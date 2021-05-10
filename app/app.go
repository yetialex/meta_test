package app

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/yetialex/meta_test/app/handlers"
	"github.com/yetialex/meta_test/app/migrate"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type App struct {
	logger   *zap.Logger
	pool     *pgxpool.Pool
	config   *Config
	Handlers *handlers.Handlers
}

func (app *App) Close() {
}

func NewApp(logger *zap.Logger, config *Config, pool *pgxpool.Pool) (*App, error) {
	return &App{
		logger:   logger,
		pool:     pool,
		config:   config,
		Handlers: handlers.NewHandlers(logger, pool),
	}, nil
}

func (app *App) MigrateDB() {
	conn, err := app.pool.Acquire(context.Background())
	defer conn.Release()
	if err != nil {
		app.logger.Fatal("Unable to acquire a database connection", zap.Error(err))
	}
	app.migrateDatabase(conn.Conn())
}

func (app *App) migrateDatabase(conn *pgx.Conn) {
	migrator, err := migrate.NewMigrator(conn, "schema_version")
	if err != nil {
		app.logger.Fatal("Unable to create a migrator", zap.Error(err))
	}

	err = migrator.LoadMigrations("./migrations")
	if err != nil {
		app.logger.Fatal("Unable to load migrations", zap.Error(err))
	}

	err = migrator.Migrate(func(err error) (retry bool) {
		app.logger.Error("Commit failed during migration, retrying", zap.Error(err))
		return true
	})

	if err != nil {
		app.logger.Fatal("Unable to migrate", zap.Error(err))
	}

	ver, err := migrator.GetCurrentVersion()
	if err != nil {
		app.logger.Fatal("Unable to get current schema version", zap.Error(err))
	}

	app.logger.Info("Migration done", zap.Int32("Current schema version", ver))
}
