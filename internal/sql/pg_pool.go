package sql

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreatePool(ctx context.Context, connectionString string) *pgxpool.Pool {
	pool, err := pgxpool.New(ctx, connectionString)
	if err != nil {
		log.Fatalf("Unable to connect to postgresql: %v\n", err)
	}

	conn, err := pool.Acquire(ctx)
	if err != nil {
		log.Fatalf("Unable to aquire db connection: %v\n", err)
	}
	defer conn.Release()

	var version string
	conn.QueryRow(ctx, `SELECT version()`).Scan(&version)

	log.Printf("Postgres version: %s\n", version)

	return pool
}
