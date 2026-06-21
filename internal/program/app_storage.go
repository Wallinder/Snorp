package program

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage interface {
}

type FileStorage struct {
}

type Postgres struct {
	PgPool *pgxpool.Pool
}

func (app *Application) NewStorage(ctx context.Context, kind string) (Storage, error) {
	switch kind {
	case "postgres":
		pool, err := app.Config.Postgres.NewConnectionPool(ctx)
		if err != nil {
			return nil, err
		}
		return &Postgres{PgPool: pool}, nil
	case "file":
		return &FileStorage{}, nil
	default:
		return nil, fmt.Errorf("unknown storage type: %s", kind)
	}
}
