package app

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/kloudlite/api/apps/console/internal/domain"
	"github.com/kloudlite/api/apps/console/internal/entities"
	msgOfficeT "github.com/kloudlite/api/apps/message-office/types"
	"github.com/kloudlite/api/pkg/errors"
	fn "github.com/kloudlite/api/pkg/functions"
	"github.com/kloudlite/api/pkg/logging"
	"github.com/kloudlite/api/pkg/messaging"
	msgTypes "github.com/kloudlite/api/pkg/messaging/types"
	t "github.com/kloudlite/api/pkg/types"
	"github.com/kloudlite/operator/operators/resource-watcher/types"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type ResourceUpdateConsumer messaging.Consumer

func newResourceContext(ctx domain.ConsoleContext, environmentName string) domain.ResourceContext {
	return domain.ResourceContext{
		ConsoleContext:  ctx,
		EnvironmentName: environmentName,
	}
}

var (
	appsGVK                  = fn.GVK("crds.kloudlite.io/v1", "App")
	externalAppsGVK          = fn.GVK("crds.kloudlite.io/v1", "ExternalApp")
	environmentGVK           = fn.GVK("crds.kloudlite.io/v1", "Environment")
	deviceGVK                = fn.GVK("wireguard.kloudlite.io/v1", "Device")
	configGVK                = fn.GVK("v1", "ConfigMap")
	secretGVK                = fn.GVK("v1", "Secret")
	routerGVK                = fn.GVK("crds.kloudlite.io/v1", "Router")
	managedResourceGVK       = fn.GVK("crds.kloudlite.io/v1", "ManagedResource")
	projectManagedServiceGVK = fn.GVK("crds.kloudlite.io/v1", "ProjectManagedService")
)

func ProcessResourceUpdates(consumer ResourceUpdateConsumer, d domain.Domain, logger logging.Logger) {
	counter := 0

	getResourceContext := func(ctx domain.ConsoleContext, rt entities.ResourceType, clusterName string, obj unstructured.Unstructured) (domain.ResourceContext, error) {
		mapping, err := d.GetEnvironmentResourceMapping(ctx, rt, clusterName, obj.GetNamespace(), obj.GetName())
		if err != nil {
			return domain.ResourceContext{}, err
		}
		if mapping == nil {
			return domain.ResourceContext{}, errors.Newf("mapping not found for %s %s/%s", rt, obj.GetNamespace(), obj.GetName())
		}

		return newResourceContext(ctx, mapping.EnvironmentName), nil
	}

	msgReader := func(msg *msgTypes.ConsumeMsg) error {
		logger := logger.WithKV("subject", msg.Subject)

		counter += 1
		logger.Debugf("[%d] received message", counter)

		ru, err := msgOfficeT.UnmarshalResourceUpdate(msg.Payload)
		if err != nil {
			logger.Errorf(err, "unmarshaling resource update")
			return nil
		}

		var rwu types.ResourceUpdate
		if err := json.Unmarshal(ru.WatcherUpdate, &rwu); err != nil {
			logger.Errorf(err, "unmarshaling into resource watcher update")
			return nil
		}

		if rwu.Object == nil {
			logger.Infof("msg.Object is nil, so could not extract any info from message, ignoring ...")
			return nil
		}

		obj := unstructured.Unstructured{Object: rwu.Object}

		mLogger := logger.WithKV(
			"gvk", obj.GetObjectKind().GroupVersionKind(),
			"resource", fmt.Sprintf("%s/%s", obj.GetNamespace(), obj.GetName()),
			"accountName", ru.AccountName,
			"clusterName", ru.ClusterName,
		)

		mLogger.Infof("received message")
		defer func() {
			mLogger.Infof("processed message")
		}()

		if len(strings.TrimSpace(ru.AccountName)) == 0 {
			logger.Infof("message does not contain 'accountName', so won't be able to find a resource uniquely, thus ignoring ...")
			return nil
		}

		if len(strings.TrimSpace(ru.ClusterName)) == 0 {
			logger.Infof("message does not contain 'clusterName', so won't be able to find a resource uniquely, thus ignoring ...")
			return nil
		}

		dctx := domain.NewConsoleContext(context.TODO(), "sys-user:console-resource-updater", ru.AccountName)

		gvkStr := obj.GetObjectKind().GroupVersionKind().String()

		resStatus, err := func() (types.ResourceStatus, error) {
			v, ok := rwu.Object[types.ResourceStatusKey]
			if !ok {
				return "", errors.NewE(fmt.Errorf("field %s not found in object", types.ResourceStatusKey))
			}
			s, ok := v.(string)
			if !ok {
				return "", errors.NewE(fmt.Errorf("field value %v is not a string", v))
			}

			return types.ResourceStatus(s), nil
		}()
		if err != nil {
			return err
		}

		opts := domain.UpdateAndDeleteOpts{MessageTimestamp: msg.Timestamp}

		switch gvkStr {
		case deviceGVK.String():
			{

				dev, err := fn.JsonConvert[entities.ConsoleVPNDevice](rwu.Object)
				if err != nil {
					return errors.NewE(err)
				}

				if v, ok := rwu.Object[types.KeyVPNDeviceConfig]; ok {
					b, err := json.Marshal(v)
					if err != nil {
						return errors.NewE(err)
					}
					var encodedStr t.EncodedString
					if err := json.Unmarshal(b, &encodedStr); err != nil {
						return errors.NewE(err)
					}
					dev.WireguardConfig = encodedStr
				}

				if resStatus == types.ResourceStatusDeleted {
					return d.OnVPNDeviceDeleteMessage(dctx, dev)
				}

				return d.OnVPNDeviceUpdateMessage(dctx, dev, resStatus, opts, ru.ClusterName)
			}
		//case projectGVK.String():
		//	{
		//		var p entities.Project
		//		if err := fn.JsonConversion(ru.Object, &p); err != nil {
		//			return errors.NewE(err)
		//		}
		//
		//		if resStatus == types.ResourceStatusDeleted {
		//			return d.OnProjectDeleteMessage(dctx, p)
		//		}
		//		return d.OnProjectUpdateMessage(dctx, p, resStatus, opts)
		//	}

		//case projectManagedServiceGVK.String():
		//	{
		//		var pmsvc entities.ProjectManagedService
		//		if err := fn.JsonConversion(ru.Object, &pmsvc); err != nil {
		//			return errors.NewE(err)
		//		}
		//
		//		mapping, err := d.GetProjectResourceMapping(dctx, entities.ResourceTypeProjectManagedService, ru.ClusterName, obj.GetNamespace(), obj.GetName())
		//		if err != nil {
		//			return err
		//		}
		//		if mapping == nil {
		//			return err
		//		}
		//
		//		if v, ok := ru.Object[types.KeyProjectManagedSvcSecret]; ok {
		//			s, err := fn.JsonConvertP[corev1.Secret](v)
		//			s.SetManagedFields(nil)
		//			if err != nil {
		//				return err
		//			}
		//			pmsvc.SyncedOutputSecretRef = s
		//		}
		//
		//		if resStatus == types.ResourceStatusDeleted {
		//			return d.OnProjectManagedServiceDeleteMessage(dctx, mapping.ProjectName, pmsvc)
		//		}
		//		return d.OnProjectManagedServiceUpdateMessage(dctx, mapping.ProjectName, pmsvc, resStatus, opts)
		//	}

		case environmentGVK.String():
			{
				var ws entities.Environment
				if err := fn.JsonConversion(rwu.Object, &ws); err != nil {
					return errors.NewE(err)
				}

				if resStatus == types.ResourceStatusDeleted {
					return d.OnEnvironmentDeleteMessage(dctx, ws)
				}
				return d.OnEnvironmentUpdateMessage(dctx, ws, resStatus, opts)
			}
		case appsGVK.String():
			{
				var app entities.App
				if err := fn.JsonConversion(rwu.Object, &app); err != nil {
					return errors.NewE(err)
				}

				rctx, err := getResourceContext(dctx, entities.ResourceTypeApp, ru.ClusterName, obj)
				if err != nil {
					return errors.NewE(err)
				}

				if resStatus == types.ResourceStatusDeleted {
					return d.OnAppDeleteMessage(rctx, app)
				}
				return d.OnAppUpdateMessage(rctx, app, resStatus, opts)
			}
		case externalAppsGVK.String():
			{
				var extApp entities.ExternalApp
				if err := fn.JsonConversion(rwu.Object, &extApp); err != nil {
					return errors.NewE(err)
				}

				rctx, err := getResourceContext(dctx, entities.ResourceTypeExternalApp, ru.ClusterName, obj)
				if err != nil {
					return errors.NewE(err)
				}

				if resStatus == types.ResourceStatusDeleted {
					return d.OnExternalAppDeleteMessage(rctx, extApp)
				}
				return d.OnExternalAppUpdateMessage(rctx, extApp, resStatus, opts)
			}
		case configGVK.String():
			{
				var config entities.Config
				if err := fn.JsonConversion(rwu.Object, &config); err != nil {
					return errors.NewE(err)
				}

				rctx, err := getResourceContext(dctx, entities.ResourceTypeConfig, ru.ClusterName, obj)
				if err != nil {
					return errors.NewE(err)
				}

				if resStatus == types.ResourceStatusDeleted {
					return d.OnConfigDeleteMessage(rctx, config)
				}
				return d.OnConfigUpdateMessage(rctx, config, resStatus, opts)
			}
		case secretGVK.String():
			{
				var secret entities.Secret
				if err := fn.JsonConversion(rwu.Object, &secret); err != nil {
					return errors.NewE(err)
				}

				if secret.Type == corev1.SecretTypeDockerConfigJson {
					// secret is an image pull secret
					ips := entities.ImagePullSecret{
						ObjectMeta: secret.ObjectMeta,
					}
					if resStatus == types.ResourceStatusDeleted {
						return d.OnImagePullSecretDeleteMessage(dctx, ips)
					}
					return d.OnImagePullSecretUpdateMessage(dctx, ips, resStatus, opts)
				}

				rctx, err := getResourceContext(dctx, entities.ResourceTypeSecret, ru.ClusterName, obj)
				if err != nil {
					return errors.NewE(err)
				}

				if resStatus == types.ResourceStatusDeleted {
					return d.OnSecretDeleteMessage(rctx, secret)
				}
				return d.OnSecretUpdateMessage(rctx, secret, resStatus, opts)
			}

		case routerGVK.String():
			{
				var router entities.Router
				if err := fn.JsonConversion(rwu.Object, &router); err != nil {
					return errors.NewE(err)
				}

				rctx, err := getResourceContext(dctx, entities.ResourceTypeRouter, ru.ClusterName, obj)
				if err != nil {
					return errors.NewE(err)
				}

				if resStatus == types.ResourceStatusDeleted {
					return d.OnRouterDeleteMessage(rctx, router)
				}
				return d.OnRouterUpdateMessage(rctx, router, resStatus, opts)
			}
		case managedResourceGVK.String():
			{
				var mres entities.ManagedResource
				if err := fn.JsonConversion(rwu.Object, &mres); err != nil {
					return errors.NewE(err)
				}

				//rctx, err := getResourceContext(dctx, entities.ResourceTypeManagedResource, ru.ClusterName, obj)
				//if err != nil {
				//	return errors.NewE(err)
				//}

				if v, ok := rwu.Object[types.KeyManagedResSecret]; ok {
					s, err := fn.JsonConvertP[corev1.Secret](v)
					if err != nil {
						mLogger.Infof("managed resource, invalid output secret received")
						return errors.NewE(err)
					}
					s.SetManagedFields(nil)
					mLogger.Infof("seting managed resource output secret")
					mres.SyncedOutputSecretRef = s
				}

				if resStatus == types.ResourceStatusDeleted {
					return d.OnManagedResourceDeleteMessage(dctx, mres.ManagedResource.Spec.ResourceTemplate.MsvcRef.Name, mres)
				}
				return d.OnManagedResourceUpdateMessage(dctx, mres.ManagedResource.Spec.ResourceTemplate.MsvcRef.Name, mres, resStatus, opts)
			}

		}
		return nil
	}

	if err := consumer.Consume(msgReader, msgTypes.ConsumeOpts{
		OnError: func(err error) error {
			logger.Errorf(err, "received while reading messages, ignoring it")
			return nil
		},
	}); err != nil {
		logger.Errorf(err, "error while consuming messages")
	}
}
