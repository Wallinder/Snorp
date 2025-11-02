package sql

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDatabase(ctx context.Context, pool *pgxpool.Pool) error {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(ctx,
		`CREATE TABLE IF NOT EXISTS archived_messages (
	      	id VARCHAR(64) PRIMARY KEY,
			type INT,
			author_id VARCHAR(64),
			global_name VARCHAR(32),
			username VARCHAR(32),
			content TEXT,
			timestamp DATE
	    )`,
	)
	if err != nil {
		return err
	}

	return nil
}
