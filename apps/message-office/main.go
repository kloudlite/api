package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"go.uber.org/fx"
	"kloudlite.io/pkg/logging"

	env "kloudlite.io/apps/message-office/internal/env"
	"kloudlite.io/apps/message-office/internal/framework"
)

func main() {
	var isDev bool
	flag.BoolVar(&isDev, "dev", false, "--dev")
	flag.Parse()

	app := fx.New(
		fx.NopLogger,

		fx.Provide(func() *env.Env {
			return env.LoadEnvOrDie()
		}),

		fx.Provide(
			func() (logging.Logger, error) {
				return logging.New(&logging.Options{Name: "message-office", Dev: isDev})
			},
		),
		// fn.FxErrorHandler(),
		framework.Module,
	)

	ctx, cancelFn := func() (context.Context, context.CancelFunc) {
		if isDev {
			return context.WithCancel(context.TODO())
		}
		return context.WithTimeout(context.TODO(), 5*time.Second)
	}()

	defer cancelFn()
	if err := app.Start(ctx); err != nil {
		panic(err)
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