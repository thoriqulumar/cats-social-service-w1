package delivery

import (
	"context"
	"github.com/thoriqulumar/cats-social-service-w1/internal/app/model"
)

type service interface {
	ValidateUser(ctx context.Context, user model.User) (err error)
	Register(ctx context.Context, data model.User) (user model.UserWithAccess, err error)
	Login(ctx context.Context, data model.LoginRequest) (user model.UserWithAccess, err error)

	MatchCat(ctx context.Context, match model.MatchRequest, userId int64)(err error)
	ValidationMatchCat(ctx context.Context, match model.MatchRequest, issuedId int64) (err error)
}

type Handler struct {
	service service
}

func New(service service) *Handler {
	return &Handler{
		service: service,
	}
}
