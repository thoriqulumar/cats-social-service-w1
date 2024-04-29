package service

import (
	"context"
	"github.com/thoriqulumar/cats-social-service-w1/internal/app/config"
	"github.com/thoriqulumar/cats-social-service-w1/internal/app/model"
	"go.uber.org/zap"
)

// repository interface, the used function is declared here
type repository interface {
	GetUserByEmail(ctx context.Context, email string) (user model.User, err error)
	CreateUser(ctx context.Context, data model.User) (err error)
}

type Service struct {
	cfg    *config.Config
	logger *zap.Logger
	repo   repository
}

func NewService(cfg *config.Config, logger *zap.Logger, repo repository) *Service {
	return &Service{
		cfg:    cfg,
		logger: logger,
		repo:   repo,
	}
}
