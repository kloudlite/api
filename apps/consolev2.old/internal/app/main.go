package app

import (
	// "context"
	"encoding/json"
	"fmt"
	"time"

	kldns "kloudlite.io/grpc-interfaces/kloudlite.io/rpc/dns"

	"github.com/gofiber/fiber/v2"
	fWebsocket "github.com/gofiber/websocket/v2"
	"google.golang.org/grpc"
	opcrds "kloudlite.io/apps/consolev2.old/internal/domain/op-crds"
	"kloudlite.io/common"
	"kloudlite.io/grpc-interfaces/kloudlite.io/rpc/auth"
	"kloudlite.io/grpc-interfaces/kloudlite.io/rpc/ci"
	"kloudlite.io/grpc-interfaces/kloudlite.io/rpc/console"
	"kloudlite.io/grpc-interfaces/kloudlite.io/rpc/finance"
	"kloudlite.io/grpc-interfaces/kloudlite.io/rpc/iam"
	"kloudlite.io/grpc-interfaces/kloudlite.io/rpc/jseval"
	"kloudlite.io/pkg/cache"
	"kloudlite.io/pkg/config"
	httpServer "kloudlite.io/pkg/http-server"
	"kloudlite.io/pkg/logging"
	lokiserver "kloudlite.io/pkg/loki-server"
	"kloudlite.io/pkg/redpanda"

	"go.uber.org/fx"
	"kloudlite.io/apps/consolev2.old/internal/app/graph"
	"kloudlite.io/apps/consolev2.old/internal/app/graph/generated"
	"kloudlite.io/apps/consolev2.old/internal/domain"
	"kloudlite.io/apps/consolev2.old/internal/domain/entities"
	"kloudlite.io/pkg/repos"
)

type WorkloadStatusConsumerEnv struct {
	Topic         string `env:"KAFKA_WORKLOAD_STATUS_TOPIC"`
	KafkaUsername string `env:"KAFKA_USERNAME"`
	KafkaPassword string `env:"KAFKA_PASSWORD"`
}

func (i *WorkloadStatusConsumerEnv) GetKafkaSASLAuth() *redpanda.KafkaSASLAuth {
	return &redpanda.KafkaSASLAuth{
		SASLMechanism: redpanda.ScramSHA256,
		User:          i.KafkaUsername,
		Password:      i.KafkaPassword,
	}
}

func (i *WorkloadStatusConsumerEnv) GetSubscriptionTopics() []string {
	return []string{
		i.Topic,
	}
}

func (i *WorkloadStatusConsumerEnv) GetConsumerGroupId() string {
	return "console-workload-consumer-2"
}

type Env struct {
	KafkaConsumerGroupId string `env:"KAFKA_GROUP_ID"`
	CookieDomain         string `env:"COOKIE_DOMAIN"`
	AuthRedisPrefix      string `env:"REDIS_AUTH_PREFIX"`
}

type InfraClientConnection *grpc.ClientConn
type IAMClientConnection *grpc.ClientConn
type JSEvalClientConnection *grpc.ClientConn
type AuthClientConnection *grpc.ClientConn
type CIClientConnection *grpc.ClientConn
type FinanceClientConnection *grpc.ClientConn
type DNSClientConnection *grpc.ClientConn

type AuthCacheClient cache.Client
type CacheClient cache.Client

var Module = fx.Module(
	"app",

	// Configs
	config.EnvFx[Env](),
	config.EnvFx[PrometheusOpts](),

	// Repos
	repos.NewFxMongoRepo[*entities.ResInstance]("res_instances", "ins", entities.ResourceIndexes),
	repos.NewFxMongoRepo[*entities.Environment]("environments", "env", entities.EnvironmentIndexes),
	repos.NewFxMongoRepo[*entities.Cluster]("clusters", "clus", entities.ClusterIndexes),
	repos.NewFxMongoRepo[*entities.EdgeRegion]("regions", "reg", entities.EdgeRegionIndexes),
	repos.NewFxMongoRepo[*entities.CloudProvider]("providers", "cp", entities.CloudProviderIndexes),
	repos.NewFxMongoRepo[*entities.Device]("devices", "dev", entities.DeviceIndexes),
	repos.NewFxMongoRepo[*entities.Project]("projects", "proj", entities.ProjectIndexes),
	repos.NewFxMongoRepo[*entities.Config]("configs", "cfg", entities.ConfigIndexes),
	repos.NewFxMongoRepo[*entities.Secret]("secrets", "sec", entities.SecretIndexes),
	repos.NewFxMongoRepo[*entities.Router]("routers", "route", entities.RouterIndexes),
	repos.NewFxMongoRepo[*entities.ManagedService]("managed_services", "mgsvc", entities.ManagedServiceIndexes),
	repos.NewFxMongoRepo[*entities.App]("apps", "app", entities.AppIndexes),
	repos.NewFxMongoRepo[*entities.ManagedResource]("managed_resources", "mgres", entities.ManagedResourceIndexes),

	// Grpc Clients

	fx.Provide(
		func(conn JSEvalClientConnection) jseval.JSEvalClient {
			return jseval.NewJSEvalClient((*grpc.ClientConn)(conn))
		},
	),

	fx.Provide(
		func(conn CIClientConnection) ci.CIClient {
			return ci.NewCIClient((*grpc.ClientConn)(conn))
		},
	),

	fx.Provide(
		func(conn IAMClientConnection) iam.IAMClient {
			return iam.NewIAMClient((*grpc.ClientConn)(conn))
		},
	),

	fx.Provide(
		func(conn AuthClientConnection) auth.AuthClient {
			return auth.NewAuthClient((*grpc.ClientConn)(conn))
		},
	),

	fx.Provide(
		func(conn DNSClientConnection) kldns.DNSClient {
			return kldns.NewDNSClient((*grpc.ClientConn)(conn))
		},
	),

	fx.Provide(
		func(conn FinanceClientConnection) finance.FinanceClient {
			return finance.NewFinanceClient((*grpc.ClientConn)(conn))
		},
	),

	// Grpc Server
	fx.Provide(fxConsoleGrpcServer),
	fx.Invoke(
		func(server *grpc.Server, consoleServer console.ConsoleServer) {
			console.RegisterConsoleServer(server, consoleServer)
		},
	),

	// Common Producer
	redpanda.NewProducerFx[redpanda.Client](),

	// Workload Message Producer
	fx.Provide(fxWorkloadMessenger),
	config.EnvFx[WorkloadStatusConsumerEnv](),
	redpanda.NewConsumerFx[*WorkloadStatusConsumerEnv](),
	fx.Invoke(
		func(domain domain.Domain, consumer redpanda.Consumer) {
			consumer.StartConsuming(
				func(msg []byte, timestamp time.Time, offset int64) error {
					var update opcrds.StatusUpdate
					if err := json.Unmarshal(msg, &update); err != nil {
						fmt.Println(err)
						return err
					}
					fmt.Println("processing", offset, string(msg), timestamp)
					if update.Stage == "EXISTS" {
						switch update.Metadata.GroupVersionKind.Kind {

						case "App",
							"Lambda",
							"Config",
							"Secret",
							"Router",
							"ManagedResource",
							"ManagedService":
							// domain.OnUpdateInstance(context.TODO(), &update)

						// case "App":
						// 	domain.OnUpdateApp(context.TODO(), &update)
						// case "Lambda":
						// 	domain.OnUpdateApp(context.TODO(), &update)
						// case "Router":
						// 	domain.OnUpdateRouter(context.TODO(), &update)
						// case "ManagedResource":
						// 	domain.OnUpdateManagedRes(context.TODO(), &update)
						// case "ManagedService":
						// 	domain.OnUpdateManagedSvc(context.TODO(), &update)

						case "Env":
							// domain.OnUpdateEnv(context.TODO(), &update)

						case "Project":
							// domain.OnUpdateProject(context.TODO(), &update)
						case "CloudProvider":
							// domain.OnUpdateProvider(context.TODO(), &update)
						case "Edge":
							// domain.OnUpdateEdge(context.TODO(), &update)
						case "Device":
							// domain.OnUpdateDevice(context.TODO(), &update)

						default:
							fmt.Println("Unknown Kind:", update.Metadata.GroupVersionKind.Kind)
						}
					}
					if update.Stage == "DELETED" {
						switch update.Metadata.GroupVersionKind.Kind {

						case "App",
							"Lambda",
							"Config",
							"Secret",
							"Router",
							"ManagedResource",
							"ManagedService":
							// domain.OnDeleteInstance(context.TODO(), &update)
						// case "App":
						// 	domain.OnDeleteApp(context.TODO(), &update)
						// case "Lambda":
						// 	domain.OnDeleteApp(context.TODO(), &update)
						// case "Router":
						// 	domain.OnDeleteRouter(context.TODO(), &update)
						// case "Project":
						// 	domain.OnDeleteProject(context.TODO(), &update)
						// case "ManagedService":
						// 	domain.OnDeleteManagedService(context.TODO(), &update)
						// case "ManagedResource":
						// 	domain.OnDeleteManagedResource(context.TODO(), &update)

						case "Env":
							// domain.OnDeleteEnv(context.TODO(), &update)
						case "Project":
							// domain.OnDeleteProject(context.TODO(), &update)
						case "CloudProvider":
							// domain.OnDeleteProvider(context.TODO(), &update)
						case "Edge":
							// domain.OnDeleteEdge(context.TODO(), &update)
						case "Device":
							// domain.OnDeleteDevice(context.TODO(), &update)
						default:
							fmt.Println("Unknown Kind:", update.Metadata.GroupVersionKind.Kind)
						}
					}
					return nil
				},
			)
		},
	),

	domain.Module,

	// Log Service
	fx.Invoke(
		func(logServer lokiserver.LogServer,
			financeClient finance.FinanceClient,
			client lokiserver.LokiClient, env *Env, cacheClient AuthCacheClient, d domain.Domain, logger logging.Logger) {
			var a *fiber.App
			a = logServer
			a.Use(
				httpServer.NewSessionMiddleware[*common.AuthSession](
					cacheClient,
					"hotspot-session",
					env.CookieDomain,
					common.CacheSessionPrefix,
				),
			)
			a.Get(
				"/build-logs", fWebsocket.New(
					func(conn *fWebsocket.Conn) {

						// ctx := d.GetSocketCtx(
						// 	conn,
						// 	cacheClient,
						// 	common.CookieName,
						// 	env.CookieDomain,
						// 	common.CacheSessionPrefix,
						// )

						appId := conn.Query("app_id", "")
						pipelineId := conn.Query("pipeline_id", "")
						pipelineRunId := conn.Query("pipeline_run_id", "")

						if len(appId) == 0 || len(pipelineId) == 0 || len(pipelineRunId) == 0 {
							logger.Infof("build logs require [app_id, pipeline_id, pipeline_run_id] in query params, missing required params, aborting ...")
							return
						}

						// app, err := d.GetApp(ctx, repos.ID(appId))
						// if err != nil {
						// 	fmt.Println(err)
						// 	conn.Close()
						// 	return
						// }
						//
						// project, err := d.GetProjectWithID(ctx, app.ProjectId)
						// if err != nil {
						// 	conn.Close()
						// 	return
						// }
						//
						// cluster, err := financeClient.GetAttachedCluster(
						// 	context.TODO(),
						// 	&finance.GetAttachedClusterIn{AccountId: string(project.AccountId)},
						// )

						// pipelineId, ok := app.Metadata["pipeline_id"]
						// if !ok {
						// 	fmt.Println("no pipeline_id found")
						// 	conn.Close()
						// 	return
						// }

						// Crosscheck session
						// if err := client.Tail(
						// 	cluster.ClusterId, []lokiserver.StreamSelector{
						// 		{
						// 			Key:       "namespace",
						// 			Operation: "=",
						// 			Value:     app.Namespace,
						// 		},
						// 		{
						// 			Key:       "app",
						// 			Operation: "=",
						// 			Value:     pipelineId,
						// 		},
						// 		{
						// 			Key:       "component",
						// 			Operation: "=",
						// 			Value:     pipelineRunId,
						// 		},
						// 	}, nil, nil, nil, nil, conn,
						// ); err != nil {
						// 	fmt.Println(err)
						// 	conn.Close()
						// 	return
						// }
					},
				),
			)
			a.Get(
				"/app-logs", fWebsocket.New(
					func(conn *fWebsocket.Conn) {
						// ctx := d.GetSocketCtx(
						// 	conn,
						// 	cacheClient,
						// 	common.CookieName,
						// 	env.CookieDomain,
						// 	common.CacheSessionPrefix,
						// )
						//
						// appId := conn.Query("app_id", "app_id")
						// app, err := d.GetApp(ctx, repos.ID(appId))
						// if err != nil {
						// 	fmt.Println(err)
						// 	conn.Close()
						// 	return
						// }
						//
						// project, err := d.GetProjectWithID(ctx, app.ProjectId)
						// if err != nil {
						// 	conn.Close()
						// 	return
						// }
						//
						// cluster, err := financeClient.GetAttachedCluster(
						// 	context.TODO(),
						// 	&finance.GetAttachedClusterIn{AccountId: string(project.AccountId)},
						// )
						//
						// Crosscheck session
						// err = client.Tail(
						// 	cluster.ClusterId, []lokiserver.StreamSelector{
						// 		{
						// 			Key:       "namespace",
						// 			Operation: "=",
						// 			Value:     app.Namespace,
						// 		},
						// 		{
						// 			Key:       "app",
						// 			Operation: "=",
						// 			Value:     app.ReadableId,
						// 		},
						// 	}, nil, nil, nil, nil, conn,
						// )
						//
						// if err != nil {
						// 	fmt.Println(err)
						// 	conn.Close()
						// 	return
						// }
					},
				),
			)
		},
	),

	fxWorkloadMessenger(),
	fxMetricsQuerySvc(),

	// GraphQL Service
	fx.Invoke(
		func(
			server *fiber.App,
			d domain.Domain,
			cacheClient AuthCacheClient,
			env *Env,
		) {
			schema := generated.NewExecutableSchema(
				generated.Config{Resolvers: &graph.Resolver{Domain: d}},
			)
			httpServer.SetupGQLServer(
				server,
				schema,
				httpServer.NewSessionMiddleware[*common.AuthSession](
					cacheClient,
					"hotspot-session",
					env.CookieDomain,
					common.CacheSessionPrefix,
				),
			)
		},
	),
)