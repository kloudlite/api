package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/kloudlite/operator/pkg/kubectl"
	"os"
	"time"

	"go.uber.org/fx"
	"k8s.io/client-go/rest"
	"kloudlite.io/pkg/k8s"
	"kloudlite.io/pkg/logging"

	"kloudlite.io/apps/accounts/internal/env"
	"kloudlite.io/apps/accounts/internal/framework"
)

func main() {
	var isDev bool
	flag.BoolVar(&isDev, "dev", false, "--dev")
	flag.Parse()

	logger, err := logging.New(&logging.Options{Name: "accounts", Dev: isDev})
	if err != nil {
		panic(err)
	}

	app := fx.New(
		fx.NopLogger,

		fx.Provide(func() logging.Logger {
			return logger
		}),

		fx.Provide(func() (*env.Env, error) {
			return env.LoadEnv()
		}),

		fx.Provide(func() (*rest.Config, error) {
			if isDev {
				return &rest.Config{
					Host: "localhost:8080",
				}, nil
			}

			return k8s.RestInclusterConfig()
		}),

		fx.Provide(func(config *rest.Config) (*kubectl.YAMLClient, error) {
			return kubectl.NewYAMLClient(config)
		}),

		fx.Provide(func(config *rest.Config) (k8s.ExtendedK8sClient, error) {
			return k8s.NewExtendedK8sClient(config)
		}),

		framework.Module,
	)

	ctx, cancelFunc := func() (context.Context, context.CancelFunc) {
		if isDev {
			return context.WithTimeout(context.TODO(), 20*time.Second)
		}
		return context.WithTimeout(context.TODO(), 5*time.Second)
	}()
	defer cancelFunc()

	if err := app.Start(ctx); err != nil {
		logger.Errorf(err, "error starting accounts app")
		logger.Infof("EXITING as errors encountered during startup")
		os.Exit(1)
	}

	fmt.Println(
		`
██████  ███████  █████  ██████  ██    ██ 
██   ██ ██      ██   ██ ██   ██  ██  ██  
██████  █████   ███████ ██   ██   ████   
██   ██ ██      ██   ██ ██   ██    ██    
██   ██ ███████ ██   ██ ██████     ██    
	`,
	)

	<-app.Done()
}
