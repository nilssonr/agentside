package postgres

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nilssonr/agentside/repository/postgres/sqlc"
)

func Dial() *pgxpool.Pool {
	pool, err := pgxpool.New(context.TODO(), os.Getenv("DATABASE_URI"))
	if err != nil {
		panic(err)
	}

	return pool
}

func Queries(db *pgxpool.Pool) *sqlc.Queries {
	return sqlc.New(db)
}
