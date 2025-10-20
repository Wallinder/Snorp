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
		`CREATE TABLE IF NOT EXISTS guilds (
            id VARCHAR(64) PRIMARY KEY,
            name VARCHAR(32),
			owner_id VARCHAR(64)
        )`,
	)
	if err != nil {
		return err
	}

	_, err = conn.Exec(ctx,
		`CREATE TABLE IF NOT EXISTS channels (
            id VARCHAR(64) PRIMARY KEY,
			guild_id VARCHAR(64),
			parent_id VARCHAR(64),
            name VARCHAR(32),
			type INT,
			topic TEXT,
			FOREIGN KEY (guild_id) REFERENCES guilds(id) ON DELETE CASCADE
        )`,
	)
	if err != nil {
		return err
	}

	return nil
}
