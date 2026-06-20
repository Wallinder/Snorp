package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	Enabled           bool   `json:"enabled"`
	ConnectionString  string `json:"connection_string"`
	MaxConns          int32  `json:"max_conns"`
	MinConns          int32  `json:"min_conns"`
	MaxConnLifetime   int32  `json:"max_conn_lifetime"`
	MaxConnIdleTime   int32  `json:"max_conn_idle_time"`
	HealthcheckPeriod int32  `json:"healthcheck_period"`
}

func (p *Postgres) NewConnectionPool(ctx context.Context) (*pgxpool.Pool, error) {
	if !p.Enabled {
		return nil, nil
	}
	config, err := pgxpool.ParseConfig(p.ConnectionString)
	if err != nil {
		return nil, err
	}
	p.setConnOpts(config)

	return pgxpool.NewWithConfig(ctx, config)
}

func (p *Postgres) setConnOpts(config *pgxpool.Config) {
	config.MaxConns = p.MaxConns
	config.MinConns = p.MinConns
	config.MaxConnLifetime = time.Duration(p.MaxConnLifetime) * time.Second
	config.MaxConnIdleTime = time.Duration(p.MaxConnIdleTime) * time.Second
	config.HealthCheckPeriod = time.Duration(p.HealthcheckPeriod) * time.Second
}
