package sql

import (
	"context"
	"log"
	"snorp/internal/api"
	"time"

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

	query := `INSERT INTO archived_messages (
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

func GetJobTimestamp(ctx context.Context, pool *pgxpool.Pool, name string) (time.Time, error) {
	var timestamp time.Time

	conn, err := pool.Acquire(ctx)
	if err != nil {
		return timestamp, err
	}
	defer conn.Release()

	query := `SELECT timestamp FROM jobs WHERE name = $1`

	err = conn.QueryRow(ctx, query, name).Scan(timestamp)
	if err != nil {
		return timestamp, nil
	}

	return timestamp, err
}

func SaveJobTimestamp(ctx context.Context, pool *pgxpool.Pool, name string, timestamp time.Time) error {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	args := pgx.NamedArgs{
		"name":      name,
		"timestamp": timestamp,
	}

	query := `INSERT INTO jobs(name, timestamp)
		VALUES (
			@name, @timestamp
		) 
		ON CONFLICT (name) DO UPDATE SET 
			timestamp = @timestamp`

	_, err = conn.Exec(ctx, query, args)
	if err != nil {
		return err
	}

	return nil
}
