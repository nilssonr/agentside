package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nilssonr/agentside/repository/postgres/sqlc"
)

func Dial(connString string) (*pgxpool.Pool, error) {
	c, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	return pgxpool.NewWithConfig(context.TODO(), c)
}

func Queries(db *pgxpool.Pool) *sqlc.Queries {
	return sqlc.New(db)
}
