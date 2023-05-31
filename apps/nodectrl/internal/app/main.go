package app

import (
	"fmt"
	"os"

	"go.uber.org/fx"
	"kloudlite.io/apps/nodectrl/internal/domain"
)

var Module = fx.Module("app",
	domain.Module,
	fx.Invoke(
		func(d domain.Domain, shutdowner fx.Shutdowner) {
			if err := d.StartJob(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			} else {
				shutdowner.Shutdown()
			}
		},
	),
)
