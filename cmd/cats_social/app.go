package main

import (
	"context"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/thoriqulumar/cats-social-service-w1/internal/app/config"
	"go.uber.org/zap"
)

type resource struct {
	DB *sqlx.DB
}

func initResources(ctx context.Context, cfg *config.Config, logger *zap.Logger) (res resource, err error) {
	// Open a database connection
	db, err := sqlx.Open("postgres", cfg.DB.ConnectionURL())
	if err != nil {
		logger.Error("error opening database", zap.Error(err))
		return
	}

	err = db.Ping()
	if err != nil {
		logger.Error("error ping the database", zap.Error(err))
		return
	}

	res.DB = db
	return
}
