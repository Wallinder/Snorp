package sql

import (
	"context"
	"log"
	"snorp/internal/api"

	"github.com/jackc/pgx/v5"
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

func InsertGlobalCommand(ctx context.Context, pool *pgxpool.Pool, command *api.ApplicationCommand) error {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	args := pgx.NamedArgs{
		"id":          command.ID,
		"name":        command.Name,
		"type":        command.Type,
		"description": command.Description,
	}

	query := `INSERT INTO global_commands (
		id, name, type, description) 
	VALUES (
		@id, @name, @type, @description
	)
	ON CONFLICT (id) DO UPDATE SET
		id = @id,
		name = @name,
		type = @type,
		description = @description`

	_, err = conn.Exec(ctx, query, args)
	if err != nil {
		return err
	}

	return nil
}

func InsertMessage(ctx context.Context, pool *pgxpool.Pool, message api.Message) error {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	args := pgx.NamedArgs{
		"id":          message.ID,
		"type":        message.Type,
		"author_id":   message.Author.ID,
		"global_name": message.Author.GlobalName,
		"username":    message.Author.Username,
		"content":     message.Content,
		"timestamp":   message.Timestamp,
	}

	query := `INSERT INTO messages (
			id, type, author_id, global_name, username, content, timestamp) 
		VALUES (
			@id, @type, @author_id, @global_name, @username, @content, @timestamp
		)`

	_, err = conn.Exec(ctx, query, args)
	if err != nil {
		return err
	}

	return nil
}
