package delivery

import (
	"context"
	"database/sql"

	"github.com/thoriqulumar/cats-social-service-w1/internal/app/model"
)

type service interface {
	ValidateUser(ctx context.Context, user model.User) (err error)
	Register(ctx context.Context, data model.User) (user model.UserWithAccess, err error)
	Login(ctx context.Context, data model.LoginRequest) (user model.UserWithAccess, err error)

	RegisterCat(ctx context.Context, data model.Cat, userId int64) (model.Cat, error)
	ValidateCat(ctx context.Context, cat model.Cat) (err error)
	GetCat(ctx context.Context, req model.GetCatRequest, userId int64) ([]model.Cat, error)
	PutCat(ctx context.Context, catReq model.PostCatRequest, catId int64) (sql.Result, error)
	ValidatePutCat(ctx context.Context, catReq model.PostCatRequest, catId int64, issuerId int64) error
	DeleteCat(ctx context.Context, id int64) (err error)
	ValidateDeleteCat(ctx context.Context, id, issuedId int64) (err error)

	MatchCat(ctx context.Context, match model.MatchRequest, issuedId int64) (data model.Match, err error)
	ValidateMatchCat(ctx context.Context, match model.MatchRequest, issuedId int64) (err error)
	DeleteMatch(ctx context.Context, id int64) (err error)
	ValidateDeleteMatchId(ctx context.Context, id, issuedId int64) (err error)
	ApproveMatch(ctx context.Context, id int64, receiverID int64) (matchID string, err error)
	RejectMatch(ctx context.Context, id int64) (matchID string, err error)
	GetMatchData(ctx context.Context, id int64) (listMatch []model.MatchData, err error)
}

type Handler struct {
	service service
}

func New(service service) *Handler {
	return &Handler{
		service: service,
	}
}
