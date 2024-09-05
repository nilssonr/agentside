package cmd

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	v1 "github.com/nilssonr/agentside/api/v1"
	"github.com/nilssonr/agentside/api/v1/middleware"
	"github.com/nilssonr/agentside/internal/queue"
	"github.com/nilssonr/agentside/internal/repository/postgres"
	"github.com/nilssonr/agentside/internal/skill"
	"github.com/nilssonr/agentside/internal/tenant"
	"github.com/nilssonr/agentside/internal/user"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the Agentside server",
	Run: func(cmd *cobra.Command, args []string) {
		if err := godotenv.Load(); err != nil {
			log.Fatal(err)
		}

		pool, err := pgxpool.New(context.TODO(), os.Getenv("DATABASE_URI"))
		if err != nil {
			log.Fatal(err)
		}

		var (
			db                  = postgres.New(pool)
			tenantRepository    = postgres.NewTenantRepository(db)
			tenantService       = tenant.NewService(tenantRepository)
			userRepository      = postgres.NewUserRepository(db)
			userService         = user.NewService(userRepository)
			userSkillRepository = postgres.NewUserSkillRepository(db)
			userSkillService    = user.NewSkillService(userSkillRepository)
			skillRepository     = postgres.NewSkillRepository(db)
			skillService        = skill.NewService(skillRepository)
			queueRepository     = postgres.NewQueueRepository(db)
			queueService        = queue.NewService(queueRepository)
		)

		swagger, _ := v1.GetSwagger()
		swagger.Servers = nil

		router := chi.NewRouter()
		handler := v1.AgentsideHandler{
			TenantService:    tenantService,
			UserService:      userService,
			UserSkillService: userSkillService,
			SkillService:     skillService,
			QueueService:     queueService,
		}

		v1.HandlerWithOptions(handler, v1.ChiServerOptions{
			BaseURL:    "",
			BaseRouter: router,
			Middlewares: []v1.MiddlewareFunc{
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

		log.Println("Agentside API listening on", os.Getenv("HTTP_ADDR"))
		http.ListenAndServe(os.Getenv("HTTP_ADDR"), router)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
