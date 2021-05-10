package handlers

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type Handlers struct {
	Core *CoreHandlers
}

func NewHandlers(logger *zap.Logger, pool *pgxpool.Pool) *Handlers {
	return &Handlers{
		Core: NewCoreHandlers(logger, pool),
	}
}
