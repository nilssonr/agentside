package cmd

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/nilssonr/agentside/api/pb"
	"github.com/nilssonr/agentside/api/rest"
	"github.com/nilssonr/agentside/auth"
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
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type authenticationServer struct {
	pb.UnimplementedAuthenticationServer

	UserService user.Service
}

func newAuthenticationServer(us user.Service) pb.AuthenticationServer {
	return authenticationServer{
		UnimplementedAuthenticationServer: pb.UnimplementedAuthenticationServer{},
		UserService:                       us,
	}
}

func (as authenticationServer) Authenticate(ctx context.Context, request *pb.AuthenticationRequest) (*pb.User, error) {
	user, err := as.UserService.GetUserByEmailAddress(ctx, request.TenantId, request.Username)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, fmt.Errorf("no user found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return nil, err
	}

	result := &pb.User{
		Id:             user.ID,
		FirstName:      user.Firstname,
		LastName:       user.LastModifiedBy,
		EmailAddress:   user.EmailAddress,
		TenantId:       user.TenantID,
		LastModifiedAt: &timestamppb.Timestamp{Seconds: user.LastModifiedAt.Unix()},
		LastModifiedBy: user.LastModifiedBy,
	}

	return result, nil
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			httpAddr     = os.Getenv("HTTP_ADDR")
			grpcAddr     = os.Getenv("GRPC_ADDR")
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
			authClientRepository           = postgres.NewAuthClientRepository(db)
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
			authClientService           = auth.NewClientService(authClientRepository, logger)
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

		// Initialize gRPC server
		logger.Info("Starting Agentside gRPC server",
			zap.String("addr", grpcAddr))

		lis, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var opts []grpc.ServerOption

		srv := grpc.NewServer(opts...)
		pb.RegisterAuthenticationServer(srv, newAuthenticationServer(userService))
		go srv.Serve(lis)

		// Initialize HTTP Server
		logger.Info("Starting Agentside HTTP server",
			zap.String("addr", httpAddr))

		err = http.ListenAndServe(httpAddr, rest.NewRouter(
			rest.WithLogger(logger),
			rest.WithAuth0Audience(authAudience),
			rest.WithAuth0Domain(authDomain),
			rest.WithAuthClientService(authClientService),
			rest.WithCustomerService(customerService),
			rest.WithCustomerAddressService(customerAddressService),
			rest.WithCustomerEmailAddressService(customerEmailAddressService),
			rest.WithCustomerNoteService(customerNoteService),
			rest.WithCustomerPhoneNumberService(customerPhoneNumberService),
			rest.WithInteractionService(interactionService),
			rest.WithQueueService(queueService),
			rest.WithQueueSkillService(queueSkillService),
			rest.WithSkillService(skillService),
			rest.WithTenantService(tenantService),
			rest.WithUserService(userService),
			rest.WithUserSkillservice(userSkillService),
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
