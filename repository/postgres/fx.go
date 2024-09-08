package postgres

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module(
	"postgres",

	// Connection
	fx.Provide(NewPgxPool),
	fx.Provide(Queries),

	// Customers
	fx.Provide(NewCustomerRepository),
	fx.Provide(NewCustomerAddressRepository),
	fx.Provide(NewCustomerPhoneNumberRepository),
	fx.Provide(NewCustomerEmailAddressRepository),
	fx.Provide(NewCustomerNoteRepository),

	// Interactions
	fx.Provide(NewInteractionRepository),

	// Queues
	fx.Provide(NewQueueRepository),

	// Skills
	fx.Provide(NewSkillRepository),

	// Tenants
	fx.Provide(NewTenantRepository),

	// Users
	fx.Provide(NewUserRepository),
	fx.Provide(NewUserSkillRepository),
)

func NewPgxPool(lc fx.Lifecycle, logging *zap.Logger) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.TODO(), os.Getenv("DATABASE_URI"))
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logging.Info("Starting the database pool")
			return nil
		},
		OnStop: func(context.Context) error {
			logging.Info("Stopping database pool")
			return nil
		},
	})

	return pool, nil
}
