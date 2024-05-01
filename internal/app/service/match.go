package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/thoriqulumar/cats-social-service-w1/internal/app/model"
	"go.uber.org/zap"
)

func (s *Service) MatchCat(ctx context.Context, match model.MatchRequest, issuedId int64) (err error) {
	err = s.repo.MatchCat(ctx, match, issuedId)
	if err != nil {
		s.logger.Error("failed to create match", zap.Error(err))
		return
	}

	return nil
}

func (s *Service) ValidationMatchCat(ctx context.Context, match model.MatchRequest, issuedId int64) (err error) {
	// validate gender userCatId and matchCatId are not same
	userCatData, err := s.repo.GetCatByID(ctx, match.UserCatId)
	if err != nil && err == sql.ErrNoRows {
		return errors.New("userCatId is not found")
	}

	matchCatData, err := s.repo.GetCatByID(ctx, match.MatchCatId)
	if err != nil && err == sql.ErrNoRows {
		return errors.New("matchCatId is not found")
	}

	// check if userCatId is owned by userId
	_, err = s.repo.GetCatOwnerByID(ctx, match.UserCatId, issuedId)
	if err != nil && err == sql.ErrNoRows {
		return errors.New("issuedId not owner of userCatId")
	}

	// check if cat sex is not same
	if matchCatData.Sex == userCatData.Sex {
		return errors.New("Cat cannot same sex")
	}

	return nil
}

func (s *Service) DeleteMatch(ctx context.Context, id, issuedId int64) (err error){
	// check issuedId and id match
	_, err = s.repo.GetMatchByIdAndIssuedId(ctx, id, issuedId)
	if err != nil && err == sql.ErrNoRows {
		return errors.New("failed to delete, match data is not owned by this issuedId")
	}

	err = s.repo.DeleteMatchById(ctx, id)
	if err != nil {
		s.logger.Error("failed to delete match", zap.Error(err))
		return
	}

	return nil
}