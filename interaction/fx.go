package interaction

import "go.uber.org/fx"

var Module = fx.Module(
	"interaction",
	fx.Provide(NewService),
)
