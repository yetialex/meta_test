package core

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type Handlers struct {
	logger *zap.Logger
	pool   *pgxpool.Pool
}

func NewHandlers(logger *zap.Logger, pool *pgxpool.Pool) *Handlers {
	return &Handlers{
		logger: logger,
		pool:   pool,
	}
}
