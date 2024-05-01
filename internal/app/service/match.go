package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/thoriqulumar/cats-social-service-w1/internal/app/model"
	"go.uber.org/zap"
)

func (s *Service) MatchCat(ctx context.Context, match model.MatchRequest, issuedId int64) (data model.Match, err error) {
	data, err = s.repo.MatchCat(ctx, match, issuedId)
	if err != nil {
		return
	}
	
	return data, nil
}

func (s *Service) ValidationMatchCat(ctx context.Context, match model.MatchRequest, issuedId int64) (err error) {
	// validate gender userCatId and matchCatId are not same
	_, err = s.repo.GetCatByID(ctx, match.UserCatId)
	if err != nil && err == sql.ErrNoRows {
		return errors.New("userCatId is not found")
	}

	_, err = s.repo.GetCatByID(ctx, match.MatchCatId)
	if err != nil && err == sql.ErrNoRows {
		return errors.New("matchCatId is not found")
	}

	// check if userCatId is owned by userId
	_, err = s.repo.GetCatOwnerByID(ctx, match.UserCatId, issuedId)
	if err != nil {
		if err != sql.ErrNoRows {
			return errors.New("issuedId not owner of userCatId")
		}
		// Handle case where no cat was found (data is zero-value Cat)
		return errors.New("issuedId not owner of userCatId")
	}
	

	return nil
}

func (s *Service) ValidationRequestCat(ctx context.Context, match model.MatchRequest, issuedId int64) (err error) {
	// validate gender userCatId and matchCatId are not same
	userCatData, err := s.repo.GetCatByID(ctx, match.UserCatId)
	if err != nil && err == sql.ErrNoRows {
		return errors.New("userCatId is not found")
	}

	matchCatData, err := s.repo.GetCatByID(ctx, match.MatchCatId)
	if err != nil && err == sql.ErrNoRows {
		return errors.New("matchCatId is not found")
	}

	// check if cats are from same owner
	if matchCatData.OwnerId == userCatData.OwnerId {
		return errors.New("userCat and matchCat are belong to same owner")
	}

	// check if cat already matched
	if matchCatData.IsAlreadyMatched || userCatData.IsAlreadyMatched{
		return errors.New("userCat or matchCat already being matched")
	}

	// check if cat is not same sex
	if matchCatData.Sex == userCatData.Sex {
		return errors.New("Cat cannot same sex")
	}

	return nil
}

func (s *Service) DeleteMatch(ctx context.Context, id, issuedId int64) (err error){
	err = s.repo.DeleteMatchById(ctx, id)
	if err != nil {
		s.logger.Error("failed to delete match", zap.Error(err))
		return
	}

	return nil
}


func (s *Service) ValidateDeleteMatchId(ctx context.Context, id, issuedId int64) (err error){
	// check issuedId and id match
	_, err = s.repo.GetMatchByIdAndIssuedId(ctx, id, issuedId)
	if err != nil && err == sql.ErrNoRows {
		return errors.New("matchId is not found")
	}

	return nil
}

func (s *Service) ValidateMatchIsApproved(ctx context.Context, id, issuedId int64) (err error){
	// check issuedId and id match
	match, _ := s.repo.GetMatchByIdAndIssuedId(ctx, id, issuedId)
	if err != nil && err == sql.ErrNoRows {
		return
	}

	if match.IsApprovedOrRejected{
		return errors.New("matchId is already approved / reject")
	}

	return nil
}