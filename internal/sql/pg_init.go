package sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InitDatabase(ctx context.Context, conn *pgx.Conn) error {
	_, err := conn.Exec(ctx, `CREATE DATABASE IF NOT EXISTS snorp`)
	if err != nil {
		return err
	}

	_, err = conn.Exec(ctx,
		`CREATE TABLE IF NOT EXISTS guilds (
            id VARCHAR(64) PRIMARY KEY,
            name TEXT,
			owner VARCHAR(32)
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
			FOREIGN KEY guild_id REFERENCES guilds(id)
        )`,
	)
	if err != nil {
		return err
	}

	_, err = conn.Exec(ctx, // combination of guild members and guild users
		`CREATE TABLE IF NOT EXISTS users (
			id VARCHAR(64) PRIMARY KEY,
            username VARCHAR(32),
			global_name VARCHAR(32),
			primary_guild VARCHAR(64),
		)`,
	)
	if err != nil {
		return err
	}

	_, err = conn.Exec(ctx, // guild <-> users mapping
		`CREATE TABLE IF NOT EXISTS users_mapping (
			user_id VARCHAR(64),
			guild_id VARCHAR(64),
			PRIMARY KEY (user_id, guild_id),
			FOREIGN KEY (user_id) REFERENCES users(id),
			FOREIGN KEY (guild_id) REFERENCES guilds(id)
		)`,
	)
	if err != nil {
		return err
	}

	return nil
}
