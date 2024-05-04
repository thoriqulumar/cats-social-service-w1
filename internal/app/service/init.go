package service

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"

	"github.com/thoriqulumar/cats-social-service-w1/internal/app/config"
	"github.com/thoriqulumar/cats-social-service-w1/internal/app/model"
	"go.uber.org/zap"
)

// repository interface, the used function is declared here
//
//go:generate mockgen -source=init.go -destination=mocks/mock_repository.go -package=mocks
type repository interface {
	GetUserByEmail(ctx context.Context, email string) (user model.User, err error)
	GetUserById(ctx context.Context, id int64) (data model.UserResponse, err error)
	CreateUser(ctx context.Context, data model.User) (user model.User, err error)

	CreateCat(ctx context.Context, data model.Cat) (cat model.Cat, err error)
	GetCat(ctx context.Context, query string, args []interface{}) ([]model.Cat, error)
	GetCatByID(ctx context.Context, id int64) (data model.Cat, err error)
	GetCatOwnerByID(ctx context.Context, catId, ownerId int64) (data model.Cat, err error)
	PutCat(ctx context.Context, args []interface{}) (sql.Result, error)
	DeleteCatById(ctx context.Context, id int64) (err error)

	MatchCat(ctx context.Context, data model.MatchRequest, issuedId, receiverID int64) (model.Match, error)
	GetMatchByID(ctx context.Context, id int64) (data model.Match, err error)
	GetMatchByIdAndIssuedId(ctx context.Context, id, issuedId int64) (data model.Match, err error)
	DeleteMatchById(ctx context.Context, id int64) (err error)
	UpdateMatchStatus(ctx context.Context, id int64, status model.MatchStatus) (err error)
	GetMatchByUserCatIds(ctx context.Context, userCatIDs []int64) (listData []model.Match, err error)
	GetMatchByMatchCatIds(ctx context.Context, matchCatIDs []int64) (listData []model.Match, err error)
	GetAllMatchData(ctx context.Context, id int64) (list *sqlx.Rows, err error)
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
