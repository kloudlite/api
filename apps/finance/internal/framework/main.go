package framework

import (
	"fmt"
	"go.uber.org/fx"
	"kloudlite.io/apps/finance/internal/app"
	"kloudlite.io/pkg/cache"
	"kloudlite.io/pkg/config"
	rpc "kloudlite.io/pkg/grpc"
	httpServer "kloudlite.io/pkg/http-server"
	"kloudlite.io/pkg/logger"
	"kloudlite.io/pkg/repos"
)

type AuthGRPCEnv struct {
	AuthServerHost string `env:"AUTH_SERVER_HOST"`
	AuthServerPort uint16 `env:"AUTH_SERVER_PORT"`
}

func (e *AuthGRPCEnv) GetGCPServerURL() string {
	return fmt.Sprintf("%v:%v", e.AuthServerHost, e.AuthServerPort)
}

type CommsGRPCEnv struct {
	CommsServerHost string `env:"COMMS_SERVER_HOST"`
	CommsServerPort uint16 `env:"COMMS_SERVER_PORT"`
}

func (e *CommsGRPCEnv) GetGCPServerURL() string {
	return fmt.Sprintf("%v:%v", e.CommsServerHost, e.CommsServerPort)
}

type ConsoleGRPCEnv struct {
	ConsoleServerHost string `env:"CONSOLE_SERVER_HOST"`
	ConsoleServerPort uint16 `env:"CONSOLE_SERVER_PORT"`
}

func (e *ConsoleGRPCEnv) GetGCPServerURL() string {
	return fmt.Sprintf("%v:%v", e.ConsoleServerHost, e.ConsoleServerPort)
}

type CiGrpcEnv struct {
	CiServerHost string `env:"CI_SERVER_HOST" required:"true"`
	CiServerPort uint16 `env:"CI_SERVER_PORT" required:"true"`
}

func (e *CiGrpcEnv) GetGCPServerURL() string {
	return fmt.Sprintf("%v:%v", e.CiServerHost, e.CiServerPort)
}

type IAMGRPCEnv struct {
	IAMServerHost string `env:"IAM_SERVER_HOST"`
	IAMServerPort uint16 `env:"IAM_SERVER_PORT"`
}

func (e *IAMGRPCEnv) GetGCPServerURL() string {
	return fmt.Sprintf("%v:%v", e.IAMServerHost, e.IAMServerPort)
}

type Env struct {
	DBName        string `env:"MONGO_DB_NAME"`
	DBUrl         string `env:"MONGO_URI"`
	RedisHosts    string `env:"REDIS_HOSTS"`
	RedisUsername string `env:"REDIS_USERNAME"`
	RedisPassword string `env:"REDIS_PASSWORD"`
	HttpPort      uint16 `env:"PORT"`
	HttpCors      string `env:"ORIGINS"`
	GrpcPort      uint16 `env:"GRPC_PORT"`
}

func (e *Env) GetMongoConfig() (url string, dbName string) {
	return e.DBUrl, e.DBName
}

func (e *Env) RedisOptions() (hosts, username, password string) {
	return e.RedisHosts, e.RedisUsername, e.RedisPassword
}

func (e *Env) GetHttpPort() uint16 {
	return e.HttpPort
}

func (e *Env) GetHttpCors() string {
	return e.HttpCors
}

func (e *Env) GetGRPCPort() uint16 {
	return e.GrpcPort
}

var Module = fx.Module("framework",
	fx.Provide(logger.NewLogger),
	config.EnvFx[Env](),
	config.EnvFx[ConsoleGRPCEnv](),
	config.EnvFx[IAMGRPCEnv](),
	config.EnvFx[CiGrpcEnv](),
	config.EnvFx[AuthGRPCEnv](),
	config.EnvFx[CommsGRPCEnv](),
	rpc.NewGrpcServerFx[*Env](),
	rpc.NewGrpcClientFx[*CommsGRPCEnv, app.CommsClientConnection](),
	rpc.NewGrpcClientFx[*ConsoleGRPCEnv, app.ConsoleClientConnection](),
	rpc.NewGrpcClientFx[*IAMGRPCEnv, app.IAMClientConnection](),
	rpc.NewGrpcClientFx[*CiGrpcEnv, app.CIGrpcClientConn](),
	rpc.NewGrpcClientFx[*AuthGRPCEnv, app.AuthGrpcClientConn](),
	repos.NewMongoClientFx[*Env](),
	cache.NewRedisFx[*Env](),
	httpServer.NewHttpServerFx[*Env](),
	app.Module,
)
