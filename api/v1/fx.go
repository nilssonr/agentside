package v1

import (
	"context"
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
				func(h http.Handler) http.Handler {
					return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						logger.Info("wee")
						h.ServeHTTP(w, r)
					})
				},
				middleware.OpenAPI(swagger),
				middleware.Auth0(
					os.Getenv("AUTH0_DOMAIN"),
					os.Getenv("AUTH0_AUDIENCE"),
				),
			},
			ErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
				logger.Error(err.Error(), zap.Error(err))
			},
		})

		lc.Append(fx.Hook{
			OnStart: func(context.Context) error {
				addr := os.Getenv("HTTP_ADDR")
				logger.Info("Starting HTTP server", zap.String("addr", addr))
				go http.ListenAndServe(addr, router)
				return nil
			},
			OnStop: func(context.Context) error {
				logger.Info("Stopping http server")
				return nil
			},
		})
	}),
)
