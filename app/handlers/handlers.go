package handlers

import (
	"github.com/yetialex/meta_test/app/handlers/core"
	"github.com/yetialex/meta_test/app/handlers/iba"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type Handlers struct {
	Core *core.Handlers
	Iba  *iba.Handlers
}

func NewHandlers(logger *zap.Logger, pool *pgxpool.Pool) *Handlers {
	return &Handlers{
		Core: core.NewHandlers(logger, pool),
		Iba:  iba.NewHandlers(logger, pool),
	}
}
