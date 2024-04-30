package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/thoriqulumar/cats-social-service-w1/internal/app/model"
	"go.uber.org/zap"
)


func (s *Service) MatchCat(ctx context.Context, match model.MatchRequest, userId int64)(err error){
	err = s.repo.MatchCat(ctx, match, userId)
	if err != nil {
		s.logger.Error("failed to create user", zap.Error(err))
		return
	}

	return nil
}

func (s *Service) ValidateIdCat(ctx context.Context, match model.MatchRequest)(err error){
	// TODO validate userCatId and matchCatId are found or not
	_, err = s.repo.GetCatById(ctx, match.UserCatId)
	if err != nil && err != sql.ErrNoRows {
		return errors.New("userCatId is not found")
	}

	_, err = s.repo.GetCatById(ctx, match.MatchCatId)
	if err != nil && err != sql.ErrNoRows {
		return errors.New("matchCatId is not found")
	}
	
	return nil
}

func (s *Service) ValidateGenderCat(ctx context.Context, match model.MatchRequest) (err error){
	// validate gender userCatId and matchCatId are not same
	userCatData, err := s.repo.GetCatById(ctx, match.UserCatId)
	if err != nil && err != sql.ErrNoRows {
		return errors.New("userCatId is not found")
	}

	matchCatData, err := s.repo.GetCatById(ctx, match.MatchCatId)
	if err != nil && err != sql.ErrNoRows {
		return errors.New("matchCatId is not found")
	}

	if matchCatData.Sex == userCatData.Sex{
		return errors.New("Cat cannot same sex")
	}

	return nil
}