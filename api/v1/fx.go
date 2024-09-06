package v1

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/nilssonr/agentside/api/v1/middleware"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module(
	"api",
	fx.Provide(NewCustomerHandler),
	fx.Provide(NewInteractionHandler),
	fx.Provide(NewQueueHandler),
	fx.Provide(NewSkillHandler),
	fx.Provide(NewTenantHandler),
	fx.Provide(NewUserHandler),
	fx.Provide(NewServerHandler),
	fx.Invoke(func(lc fx.Lifecycle, logger *zap.Logger, handler ServerHandler) {
		swagger, _ := GetSwagger()
		swagger.Servers = nil

		router := chi.NewRouter()
		HandlerWithOptions(handler, ChiServerOptions{
			BaseURL:    "",
			BaseRouter: router,
			Middlewares: []MiddlewareFunc{
				middleware.OpenAPI(swagger),
				middleware.Auth0(
					os.Getenv("AUTH0_DOMAIN"),
					os.Getenv("AUTH0_AUDIENCE"),
				),
			},
			ErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
				log.Println(err)
			},
		})

		lc.Append(fx.Hook{
			OnStart: func(context.Context) error {
				logger.Info("wee")
				go http.ListenAndServe(os.Getenv("HTTP_ADDR"), router)
				return nil
			},
			OnStop: func(context.Context) error {
				logger.Info("oh no")
				return nil
			},
		})
	}),
)
