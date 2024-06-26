package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"os"
	"strings"
	"time"

	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/pkg/errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/yaml"

	"github.com/kloudlite/api/apps/tenant-agent/internal/env"
	proto_rpc "github.com/kloudlite/api/apps/tenant-agent/internal/proto-rpc"
	t "github.com/kloudlite/api/apps/tenant-agent/types"
	"github.com/kloudlite/operator/grpc-interfaces/grpc/messages"
	libGrpc "github.com/kloudlite/operator/pkg/grpc"
	"github.com/kloudlite/operator/pkg/kubectl"

	"github.com/kloudlite/operator/pkg/logging"
	apiErrors "k8s.io/apimachinery/pkg/api/errors"
)

type grpcHandler struct {
	inMemCounter   int64
	yamlClient     kubectl.YAMLClient
	logger         logging.Logger
	ev             *env.Env
	msgDispatchCli messages.MessageDispatchServiceClient
	isDev          bool
}

const (
	MaxConnectionDuration = 45 * time.Second
)

func (g *grpcHandler) handleErrorOnApply(ctx context.Context, err error, msg t.AgentMessage) error {
	g.logger.Debugf("[ERROR]: %s", err.Error())

	b, err := json.Marshal(t.AgentErrMessage{
		AccountName: msg.AccountName,
		ClusterName: msg.ClusterName,
		Error:       err.Error(),
		Action:      msg.Action,
		Object:      msg.Object,
	})
	if err != nil {
		return errors.NewE(err)
	}

	_, err = g.msgDispatchCli.ReceiveError(ctx, &messages.ErrorData{
		ProtocolVersion: g.ev.GrpcMessageProtocolVersion,
		Message:         b,
	})
	return err
}

func NewAuthorizedGrpcContext(ctx context.Context, accessToken string) context.Context {
	return metadata.NewOutgoingContext(ctx, metadata.Pairs("authorization", accessToken))
}

func (g *grpcHandler) handleMessage(gctx context.Context, msg t.AgentMessage) error {
	g.inMemCounter++
	ctx, cf := func() (context.Context, context.CancelFunc) {
		if g.isDev {
			return context.WithCancel(gctx)
		}
		return context.WithTimeout(gctx, 3*time.Second)
	}()
	defer cf()

	if msg.Object == nil {
		g.logger.Infof("msg.Object is nil, could not process anything out of this message, ignoring ...")
		return nil
	}

	obj := unstructured.Unstructured{Object: msg.Object}
	mLogger := g.logger.WithKV("gvk", obj.GetObjectKind().GroupVersionKind().String()).WithKV("clusterName", msg.ClusterName).WithKV("accountName", msg.AccountName).WithKV("action", msg.Action)

	mLogger.Infof("[%d] received message", g.inMemCounter)

	if len(strings.TrimSpace(msg.AccountName)) == 0 {
		return g.handleErrorOnApply(ctx, errors.Newf("field 'accountName' must be defined in message"), msg)
	}

	switch msg.Action {
	case t.ActionApply:
		{
			ann := obj.GetAnnotations()
			if ann == nil {
				ann = make(map[string]string, 2)
			}

			obj.SetAnnotations(ann)

			b, err := yaml.Marshal(msg.Object)
			if err != nil {
				return g.handleErrorOnApply(ctx, err, msg)
			}

			if _, err := g.yamlClient.ApplyYAML(ctx, b); err != nil {
				mLogger.Errorf(err, "[%d] [error-on-apply]: yaml: \n%s\n", g.inMemCounter, b)
				mLogger.Infof("[%d] failed to process message", g.inMemCounter)
				return g.handleErrorOnApply(ctx, err, msg)
			}
			mLogger.Infof("[%d] processed message", g.inMemCounter)
		}
	case t.ActionDelete:
		{
			if err := g.yamlClient.DeleteResource(ctx, &obj); err != nil {
				mLogger.Infof("[%d] [error-on-delete]: %v", g.inMemCounter, err)
				if apiErrors.IsNotFound(err) {
					mLogger.Infof("[%d] processed message, resource does not exist, might already be deleted", g.inMemCounter)
					return g.handleErrorOnApply(ctx, err, msg)
				}
				mLogger.Infof("[%d] failed to process message", g.inMemCounter)
			}
			mLogger.Infof("[%d] processed message", g.inMemCounter)
		}
	case t.ActionRestart:
		{
			if err := g.yamlClient.RolloutRestart(ctx, kubectl.Deployment, obj.GetNamespace(), obj.GetLabels()); err != nil {
				return err
			}
			mLogger.Infof("[%d] rolled out deployments", g.inMemCounter)

			if err := g.yamlClient.RolloutRestart(ctx, kubectl.StatefulSet, obj.GetNamespace(), obj.GetLabels()); err != nil {
				return err
			}

			mLogger.Infof("[%d] rolled out statefulsets", g.inMemCounter)
			mLogger.Infof("[%d] processed message", g.inMemCounter)
		}
	default:
		{
			err := errors.Newf("invalid action (%s)", msg.Action)
			mLogger.Infof("[%d] [error]: %s", err.Error())
			mLogger.Infof("[%d] failed to process message", g.inMemCounter)
			return g.handleErrorOnApply(ctx, err, msg)
		}
	}

	return nil
}

func (g *grpcHandler) ensureAccessToken() error {
	if g.ev.AccessToken == "" {
		g.logger.Infof("waiting on clusterToken exchange for accessToken")
	}

	ctx := NewAuthorizedGrpcContext(context.TODO(), g.ev.AccessToken)

	validationOut, err := g.msgDispatchCli.ValidateAccessToken(ctx, &messages.ValidateAccessTokenIn{
		ProtocolVersion: g.ev.GrpcMessageProtocolVersion,
	})
	if err != nil {
		g.logger.Errorf(err, "validating access token")
		validationOut = nil
	}

	if validationOut != nil && validationOut.Valid {
		g.logger.Infof("accessToken is valid, proceeding with it ...")
		return nil
	}

	g.logger.Infof("accessToken is invalid, requesting new accessToken ...")

	out, err := g.msgDispatchCli.GetAccessToken(ctx, &messages.GetAccessTokenIn{
		ProtocolVersion: g.ev.GrpcMessageProtocolVersion,
		ClusterToken:    g.ev.ClusterToken,
	})
	if err != nil {
		return errors.NewE(err)
	}

	g.logger.Infof("valid access token has been obtained, persisting it in k8s secret (%s/%s)...", g.ev.AccessTokenSecretNamespace, g.ev.AccessTokenSecretName)

	s, err := g.yamlClient.Client().CoreV1().Secrets(g.ev.AccessTokenSecretNamespace).Get(context.TODO(), g.ev.AccessTokenSecretName, metav1.GetOptions{})
	if err != nil {
		return errors.NewE(err)
	}

	if s.Data == nil {
		s.Data = make(map[string][]byte, 1)
	}
	s.Data["ACCESS_TOKEN"] = []byte(out.AccessToken)
	s.Data["ACCOUNT_NAME"] = []byte(out.AccountName)
	s.Data["CLUSTER_NAME"] = []byte(out.ClusterName)
	_, err = g.yamlClient.Client().CoreV1().Secrets(g.ev.AccessTokenSecretNamespace).Update(context.TODO(), s, metav1.UpdateOptions{})
	if err != nil {
		return errors.NewE(err)
	}

	g.ev.AccessToken = out.AccessToken

	if g.ev.ResourceWatcherNamespace != "" {
		// need to restart resource watcher
		d, err := g.yamlClient.Client().AppsV1().Deployments(g.ev.ResourceWatcherNamespace).Get(ctx, g.ev.ResourceWatcherName, metav1.GetOptions{})
		if err != nil {
			return errors.NewE(err)
		}
		podLabelSelector := metav1.LabelSelector{}
		for k, v := range d.Spec.Selector.MatchLabels {
			metav1.AddLabelToSelector(&podLabelSelector, k, v)
		}

		if err := g.yamlClient.Client().CoreV1().Pods(g.ev.ResourceWatcherNamespace).DeleteCollection(ctx, metav1.DeleteOptions{}, metav1.ListOptions{LabelSelector: metav1.FormatLabelSelector(&podLabelSelector)}); err != nil {
			g.logger.Errorf(err, "failed to delete pods for resource watcher")
		}
		g.logger.Infof("deleted all pods for resource watcher, they will be recreated")
	}

	return nil
}

func (g *grpcHandler) run(rctx context.Context, cf context.CancelFunc) error {
	defer cf()
	ctx := NewAuthorizedGrpcContext(rctx, g.ev.AccessToken)

	g.logger.Infof("asking message office to start sending actions")
	msgActionsCli, err := g.msgDispatchCli.SendActions(ctx, &messages.Empty{})
	if err != nil {
		return errors.NewE(err)
	}

	for {
		if err := ctx.Err(); err != nil {
			return err
		}

		var msg t.AgentMessage
		a, err := msgActionsCli.Recv()
		if err != nil {
			if status.Code(err) == codes.Unavailable {
				g.logger.Infof("server unavailable, (may be, Gateway Timed Out 504), reconnecting ...")
				return nil
			}
			if status.Code(err) == codes.DeadlineExceeded {
				g.logger.Infof("Connection Timed Out, reconnecting ...")
				return nil
			}
			return err
		}

		if err := json.Unmarshal(a.Message, &msg); err != nil {
			g.logger.Errorf(err, "[ERROR] while json unmarshal")
			return errors.NewE(err)
		}

		if err := g.handleMessage(ctx, msg); err != nil {
			g.logger.Errorf(err, "[ERROR] while handling message")
			return errors.NewE(err)
		}
	}
}

func main() {
	var isDev bool
	flag.BoolVar(&isDev, "dev", false, "--dev")
	flag.Parse()

	ev := env.GetEnvOrDie()

	logger := logging.NewOrDie(&logging.Options{Name: "kloudlite-agent", Dev: isDev})

	logger.Infof("waiting for GRPC connection to happen")

	yamlClient := func() kubectl.YAMLClient {
		if isDev {
			logger.Debugf("connecting to k8s over host addr (%s)", "localhost:8081")
			return kubectl.NewYAMLClientOrDie(&rest.Config{Host: "localhost:8081"}, kubectl.YAMLClientOpts{Logger: logger})
		}
		config, err := rest.InClusterConfig()
		if err != nil {
			panic(err)
		}
		return kubectl.NewYAMLClientOrDie(config, kubectl.YAMLClientOpts{Logger: logger})
	}()

	g := grpcHandler{
		inMemCounter: 0,
		yamlClient:   yamlClient,
		logger:       logger,
		ev:           ev,
		isDev:        isDev,
	}

	vps := &vectorGrpcProxyServer{
		realVectorClient: nil,
		logger:           logger,
		accessToken:      ev.AccessToken,
		errCh:            nil,
	}

	gs := libGrpc.NewGrpcServer(libGrpc.GrpcServerOpts{Logger: logger})
	proto_rpc.RegisterVectorServer(gs.GrpcServer, vps)

	go func() {
		err := gs.Listen(ev.VectorProxyGrpcServerAddr)
		if err != nil {
			logger.Error(err)
			os.Exit(1)
		}
	}()

	common.PrintReadyBanner()

	for {
		logger.Debugf("trying to connect to message office grpc (%s)", ev.GrpcAddr)
		cc, err := func() (*grpc.ClientConn, error) {
			// if isDev {
			// 	logger.Infof("attempting grpc connect over %s", ev.GrpcAddr)
			// 	return libGrpc.Connect(ev.GrpcAddr, libGrpc.ConnectOpts{
			// 		SecureConnect: false,
			// 		Timeout:       20 * time.Second,
			// 	})
			// }
			logger.Infof("attempting grpc connect over %s", ev.GrpcAddr)
			return libGrpc.ConnectSecure(ev.GrpcAddr)
		}()
		if err != nil {
			log.Fatalf("Failed to connect after retries: %v", err)
		}

		logger.Infof("GRPC connection to message-office (%s) successful", ev.GrpcAddr)

		g.msgDispatchCli = messages.NewMessageDispatchServiceClient(cc)

		if err := g.ensureAccessToken(); err != nil {
			logger.Errorf(err, "ensuring access token")
		}

		ctx, cf := context.WithTimeout(context.TODO(), MaxConnectionDuration)

		vps.accessToken = g.ev.AccessToken
		vps.realVectorClient = proto_rpc.NewVectorClient(cc)
		vps.errCh = make(chan error, 1)

		go func() {
			if err := g.run(ctx, cf); err != nil {
				logger.Errorf(err, "running grpc sendActions")
			}
		}()

		select {
		case err := <-vps.errCh:
			{
				logger.Errorf(err, "error from vector grpc proxy server")
				cf()
			}
		case <-ctx.Done():
			{
				logger.Debugf("run context done, reconnecting ...")
			}
		}

		if err = cc.Close(); err != nil {
			logger.Errorf(err, "Failed to close connection")
		}
	}
}
