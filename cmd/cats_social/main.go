package main

import (
	"context"

	"github.com/thoriqulumar/cats-social-service-w1/internal/app/config"
	"github.com/thoriqulumar/cats-social-service-w1/internal/app/delivery"
	"github.com/thoriqulumar/cats-social-service-w1/internal/app/repository"
	"github.com/thoriqulumar/cats-social-service-w1/internal/app/service"
	"github.com/thoriqulumar/cats-social-service-w1/internal/middleware"
	"github.com/thoriqulumar/cats-social-service-w1/internal/pkg/log"
	"github.com/thoriqulumar/cats-social-service-w1/pkg/version"
	"go.uber.org/zap/zapcore"
)

func main() {
	ctx := context.Background()
	// Load the configuration file
	cfg, err := config.Load(ctx)
	if err != nil {
		panic(err)
	}
	// init logger
	logger, err := log.New(zapcore.DebugLevel, version.ServiceID, version.Version)
	if err != nil {
		panic(err)
	}
	res, err := initResources(ctx, cfg, logger)
	if err != nil {
		panic(err)
	}
	repo := repository.NewRepo(res.DB)
	s := service.NewService(cfg, logger, repo)
	h := delivery.New(s)

	secretKey := cfg.JWTSecret
	authMiddleware := middleware.AuthMiddleware(secretKey)

	initRouter(h, authMiddleware)
}
