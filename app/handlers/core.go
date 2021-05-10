package handlers

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type CoreHandlers struct {
	logger *zap.Logger
	pool   *pgxpool.Pool
}

func NewCoreHandlers(logger *zap.Logger, pool *pgxpool.Pool) *CoreHandlers {
	return &CoreHandlers{
		logger: logger,
		pool:   pool,
	}
}
