package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/nilssonr/agentside/api"
	"github.com/nilssonr/agentside/customer"
	"github.com/nilssonr/agentside/interaction"
	"github.com/nilssonr/agentside/logging"
	"github.com/nilssonr/agentside/queue"
	"github.com/nilssonr/agentside/repository/postgres"
	"github.com/nilssonr/agentside/skill"
	"github.com/nilssonr/agentside/tenant"
	"github.com/nilssonr/agentside/user"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			httpAddr     = os.Getenv("HTTP_ADDR")
			authDomain   = os.Getenv("AUTH0_DOMAIN")
			authAudience = os.Getenv("AUTH0_AUDIENCE")
		)

		logger, err := logging.NewZapLogger()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		pool, err := postgres.Dial(os.Getenv("DATABASE_URI"))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Initialize repositories
		var (
			db                             = postgres.Queries(pool)
			customerRepository             = postgres.NewCustomerRepository(db)
			customerAddressRepository      = postgres.NewCustomerAddressRepository(db)
			customerEmailAddressRepository = postgres.NewCustomerEmailAddressRepository(db)
			customerNoteRepository         = postgres.NewCustomerNoteRepository(db)
			customerPhoneNumberRepository  = postgres.NewCustomerPhoneNumberRepository(db)
			interactionRepository          = postgres.NewInteractionRepository(db)
			queueRepository                = postgres.NewQueueRepository(db)
			queueSkillRepository           = postgres.NewQueueSkillRepository(db)
			skillRepository                = postgres.NewSkillRepository(db)
			tenantRepository               = postgres.NewTenantRepository(db)
			userRepository                 = postgres.NewUserRepository(db)
			userSkillRepository            = postgres.NewUserSkillRepository(db)
		)

		// Initialize services
		var (
			customerService             = customer.NewService(customerRepository, logger)
			customerAddressService      = customer.NewAddressService(customerAddressRepository, logger)
			customerEmailAddressService = customer.NewEmailAddressService(customerEmailAddressRepository, logger)
			customerNoteService         = customer.NewNoteService(customerNoteRepository, logger)
			customerPhoneNumberService  = customer.NewPhoneNumberService(customerPhoneNumberRepository, logger)
			interactionService          = interaction.NewService(interactionRepository, logger)
			queueService                = queue.NewService(queueRepository, logger)
			queueSkillService           = queue.NewSkillService(queueSkillRepository, logger)
			skillService                = skill.NewService(skillRepository, logger)
			tenantService               = tenant.NewService(tenantRepository, logger)
			userService                 = user.NewService(userRepository, logger)
			userSkillService            = user.NewSkillService(userSkillRepository, logger)
		)

		// Initialize HTTP Server
		logger.Info("Starting Agentside HTTP server",
			zap.String("addr", httpAddr))

		err = http.ListenAndServe(httpAddr, api.NewRouter(
			api.WithLogger(logger),
			api.WithAuth0Audience(authAudience),
			api.WithAuth0Domain(authDomain),
			api.WithCustomerService(customerService),
			api.WithCustomerAddressService(customerAddressService),
			api.WithCustomerEmailAddressService(customerEmailAddressService),
			api.WithCustomerNoteService(customerNoteService),
			api.WithCustomerPhoneNumberService(customerPhoneNumberService),
			api.WithInteractionService(interactionService),
			api.WithQueueService(queueService),
			api.WithQueueSkillService(queueSkillService),
			api.WithSkillService(skillService),
			api.WithTenantService(tenantService),
			api.WithUserService(userService),
			api.WithUserSkillservice(userSkillService),
		))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
