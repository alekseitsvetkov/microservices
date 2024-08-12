package database

import (
	"context"
	"github.com/alekseytsvetkov/microservices/apps/user/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func New(cfg *config.Config) (*pgxpool.Pool, error) {
	return pgxpool.New(context.Background(), cfg.DB.Postgres.URL)
}
