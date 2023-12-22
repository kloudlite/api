package app

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/kloudlite/api/pkg/messaging"
	msgTypes "github.com/kloudlite/api/pkg/messaging/types"

	t "github.com/kloudlite/api/apps/tenant-agent/types"
	"github.com/kloudlite/container-registry-authorizer/admin"
	dbv1 "github.com/kloudlite/operator/apis/distribution/v1"

	"github.com/kloudlite/api/apps/container-registry/internal/domain"
	"github.com/kloudlite/api/apps/container-registry/internal/domain/entities"
	"github.com/kloudlite/api/apps/container-registry/internal/env"
	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/constants"
	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/api/pkg/logging"
	"github.com/kloudlite/api/pkg/types"
	common_types "github.com/kloudlite/operator/apis/common-types"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

const (
	GithubEventHeader string = "X-Github-Event"
	GitlabEventHeader string = "X-Gitlab-Event"
)

type (
	GitWebhookConsumer messaging.Consumer
	BuildRunProducer   messaging.Producer
)

func getUniqueKey(build *entities.Build, hook *domain.GitWebhookPayload) string {
	uid := fmt.Sprint(build.Id, hook.CommitHash)

	return fmt.Sprintf("%x", md5.Sum([]byte(uid)))
}

func processGitWebhooks(ctx context.Context, d domain.Domain, consumer GitWebhookConsumer, producer BuildRunProducer, logr logging.Logger, envs *env.Env) error {
	err := consumer.Consume(func(msg *msgTypes.ConsumeMsg) error {
		logger := logr.WithName("ci-webhook")
		logger.Infof("started processing")
		defer func() {
			logger.Infof("finished processing")
		}()
		var gitHook types.GitHttpHook
		if err := json.Unmarshal(msg.Payload, &gitHook); err != nil {
			logger.Errorf(err, "could not unmarshal into *GitWebhookPayload")
			return errors.NewE(err)
		}

		hook, err := func() (*domain.GitWebhookPayload, error) {
			if gitHook.GitProvider == constants.ProviderGithub {
				return d.ParseGithubHook(gitHook.Headers[GithubEventHeader][0], gitHook.Body)
			}
			if gitHook.GitProvider == constants.ProviderGitlab {
				return d.ParseGitlabHook(gitHook.Headers[GitlabEventHeader][0], gitHook.Body)
			}
			return nil, errors.New("unknown git provider")
		}()
		if err != nil {
			if _, ok := err.(*domain.ErrEventNotSupported); ok {
				fmt.Println(gitHook.GitProvider)
				logger.Infof(err.Error())
				return nil
			}
			logger.Errorf(err, "could not extract gitHook")
			return errors.NewE(err)
		}

		logger = logger.WithKV("repo", hook.RepoUrl, "provider", hook.GitProvider, "branch", hook.GitBranch)

		logger.Infof("repo: %s, branch: %s, gitprovider: %s", hook.RepoUrl, hook.GitBranch, hook.GitProvider)

		builds, err := d.ListBuildsByGit(ctx, hook.RepoUrl, hook.GitBranch, hook.GitProvider)
		if err != nil {
			return errors.NewE(err)
		}

		var pullToken string

		switch hook.GitProvider {

		case constants.ProviderGithub:
			pullToken, err = d.GithubInstallationToken(ctx, hook.RepoUrl)
			if err != nil {
				fmt.Println(err)
				return errors.NewE(err)
			}

		case constants.ProviderGitlab:
			pullToken = ""

		default:
			fmt.Println("provider not supported", hook.GitProvider)
			return errors.Newf("provider %s not supported", hook.GitProvider)
		}

		for _, build := range builds {
			if hook.GitProvider == constants.ProviderGitlab {
				pullToken, err = d.GitlabPullToken(ctx, build.CredUser.UserId)
				if err != nil {
					errorMessage := fmt.Sprintf("could not get pull token for build, Error: %s", err.Error())
					if build.ErrorMessages == nil {
						build.ErrorMessages = make(map[string]string)
					}
					if build.ErrorMessages["access-error"] != errorMessage {
						build.ErrorMessages["access-error"] = errorMessage
						_, err := d.UpdateBuildInternal(ctx, build)
						if err != nil {
							return errors.NewE(err)
						}
					}

					continue
				} else {
					if build.ErrorMessages["access-error"] != "" {
						delete(build.ErrorMessages, "access-error")
						_, err := d.UpdateBuildInternal(ctx, build)
						if err != nil {
							return errors.NewE(err)
						}
					}
				}
			}

			// pullUrl, err := domain.BuildUrl(hook.RepoUrl, pullToken)
			// if err != nil {
			// 	logger.Errorf(err, "could not build pull url")
			// 	continue
			// }

			if pullToken == "" {
				logger.Warnf("pull token is empty")
				continue
			}

			// fmt.Println("pullUrl", len(builds), pullUrl)

			i, err := admin.GetExpirationTime(fmt.Sprintf("%d%s", 1, "d"))
			if err != nil {
				return errors.NewE(err)
			}

			token, err := admin.GenerateToken(domain.KL_ADMIN, build.Spec.AccountName, string("read_write"), i, envs.RegistrySecretKey+build.Spec.AccountName)
			if err != nil {
				logger.Errorf(err, "could not generate pull-token")
				continue
			}

			uniqueKey := getUniqueKey(build, hook)

			b, err := d.GetBuildTemplate(domain.BuildJobTemplateData{
				AccountName: build.Spec.AccountName,
				Name:        uniqueKey,
				Namespace:   envs.JobBuildNamespace,
				Labels: map[string]string{
					"kloudlite.io/build-id": string(build.Id),
					"kloudlite.io/account":  build.Spec.AccountName,
					"github.com/commit":     hook.CommitHash,
				},
				Annotations: map[string]string{
					"kloudlite.io/build-id": string(build.Id),
					"kloudlite.io/account":  build.Spec.AccountName,
					"github.com/commit":     hook.CommitHash,
					"github.com/repository": hook.RepoUrl,
					"github.com/branch":     hook.GitBranch,
					"kloudlite.io/repo":     build.Spec.Registry.Repo.Name,
					"kloudlite.io/tag":      strings.Join(build.Spec.Registry.Repo.Tags, ","),
				},
				BuildOptions: build.Spec.BuildOptions,
				Registry: dbv1.Registry{
					// Password: token,
					// Username: domain.KL_ADMIN,
					// Host:     envs.RegistryHost,
					Repo: dbv1.Repo{
						Name: build.Spec.Registry.Repo.Name,
						Tags: build.Spec.Registry.Repo.Tags,
					},
				},
				CacheKeyName: build.Spec.CacheKeyName,
				GitRepo: dbv1.GitRepo{
					Url:    hook.RepoUrl,
					Branch: hook.CommitHash,
				},
				Resource: build.Spec.Resource,
				CredentialsRef: common_types.SecretRef{
					Name:      fmt.Sprint("build-run-", uniqueKey),
					Namespace: envs.JobBuildNamespace,
				},
			})
			if err != nil {
				logger.Errorf(err, "could not get build template")
				return errors.NewE(err)
			}

			sec := corev1.Secret{
				TypeMeta: metav1.TypeMeta{
					Kind:       "Secret",
					APIVersion: "v1",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      fmt.Sprint("build-run-", uniqueKey),
					Namespace: envs.JobBuildNamespace,
					Annotations: map[string]string{
						"kloudlite.io/build-run.name": uniqueKey,
					},
				},
				StringData: map[string]string{
					"registry-admin": domain.KL_ADMIN,
					"registry-host":  envs.RegistryHost,
					"registry-token": token,
					"github-token":   pullToken,
				},
			}

			var m map[string]any
			if err := yaml.Unmarshal(b, &m); err != nil {
				return errors.NewE(err)
			}

			b1, err := json.Marshal(t.AgentMessage{
				AccountName: envs.BuildClusterAccountName,
				ClusterName: envs.BuildClusterName,
				Action:      t.ActionApply,
				Object:      m,
			})
			if err != nil {
				return errors.NewE(err)
			}

			b, err = yaml.Marshal(sec)
			if err != nil {
				return errors.NewE(err)
			}

			var m2 map[string]any
			if err := yaml.Unmarshal(b, &m2); err != nil {
				return errors.NewE(err)
			}

			b2, err := json.Marshal(t.AgentMessage{
				AccountName: envs.BuildClusterAccountName,
				ClusterName: envs.BuildClusterName,
				Action:      t.ActionApply,
				Object:      m2,
			})
			if err != nil {
				return errors.NewE(err)
			}
			topic := common.GetTenantClusterMessagingTopic(envs.BuildClusterAccountName, envs.BuildClusterName)
			if err = producer.Produce(ctx, msgTypes.ProduceMsg{
				Subject: topic,
				Payload: b2,
			}); err != nil {
				return errors.NewE(err)
			}

			if err = producer.Produce(ctx, msgTypes.ProduceMsg{
				Subject: topic,
				Payload: b1,
			}); err != nil {
				return errors.NewE(err)
			}

			logger.Infof("produced message to topic=%s", topic)

			build.Status = entities.BuildStatusQueued
			_, err = d.UpdateBuildInternal(ctx, build)
			if err != nil {
				return errors.NewE(err)
			}
		}
		return nil
	}, msgTypes.ConsumeOpts{})
	if err != nil {
		return errors.NewE(err)
	}
	return nil
}
