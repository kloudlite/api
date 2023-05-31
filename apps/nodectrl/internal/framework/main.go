package framework

import (
	"go.uber.org/fx"
	"kloudlite.io/apps/nodectrl/internal/app"
)

var Module = fx.Module("framework",
	app.Module,
)
