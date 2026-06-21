package program

import (
	"context"
	"fmt"
	"snorp/internal/storage"
)

type Storage interface {
	Init()
}

func (app *Application) NewStorage(ctx context.Context, kind string) (Storage, error) {
	switch kind {
	case "postgres":
		pool, err := app.Config.Postgres.NewConnectionPool(ctx)
		if err != nil {
			return nil, err
		}
		return &storage.Postgres{PgPool: pool}, nil
	case "file":
		return &storage.FileStorage{Path: "./storage"}, nil
	default:
		return nil, fmt.Errorf("unknown storage type: %s", kind)
	}
}
