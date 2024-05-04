package service

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/thoriqulumar/cats-social-service-w1/internal/app/model"
	cerror "github.com/thoriqulumar/cats-social-service-w1/internal/pkg/error"
	"go.uber.org/zap"
)

func (s *Service) MatchCat(ctx context.Context, match model.MatchRequest, issuedId int64) (data model.Match, err error) {
	// get receiverID
	matchCatId, _ := strconv.ParseInt(match.MatchCatId, 10, 64)
	matchCat, err := s.repo.GetCatByID(ctx, matchCatId)
	if err != nil {
		return
	}

	data, err = s.repo.MatchCat(ctx, match, issuedId, matchCat.OwnerId)
	if err != nil {
		return
	}

	return data, nil
}

func (s *Service) ValidateMatchCat(ctx context.Context, match model.MatchRequest, issuedId int64) (err error) {
	// validate gender userCatId and matchCatId are not same
	matchUserCatId, _ := strconv.ParseInt(match.UserCatId, 10, 64)
	userCatData, err := s.repo.GetCatByID(ctx, matchUserCatId)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return cerror.New(http.StatusNotFound, "userCatId is not found")
	}

	matchCatId, _ := strconv.ParseInt(match.MatchCatId, 10, 64)
	matchCatData, err := s.repo.GetCatByID(ctx, matchCatId)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return cerror.New(http.StatusNotFound, "matchCatId is not found")
	}

	// check if userCatId is owned by userId
	//_, err = s.repo.GetCatOwnerByID(ctx, matchUserCatId, issuedId)
	//if err != nil {
	//	if !errors.Is(err, sql.ErrNoRows) {
	//		return cerror.New(http.StatusNotFound, "issuedId not owner of userCatId")
	//	}
	//	// Handle case where no cat was found (data is zero-value Cat)
	//	return cerror.New(http.StatusNotFound, "issuedId not owner of userCatId")
	//}

	// check if cats are from same owner
	if matchCatData.OwnerId == userCatData.OwnerId {
		return cerror.New(http.StatusBadRequest, "userCat and matchCat are belong to same owner")
	}

	// check if cat already matched
	if matchCatData.HasMatched || userCatData.HasMatched {
		return cerror.New(http.StatusBadRequest, "userCat or matchCat already being matched")
	}

	// check if cat is not same sex
	if matchCatData.Sex == userCatData.Sex {
		return cerror.New(http.StatusBadRequest, "Cat cannot same sex")
	}

	return nil
}

func (s *Service) DeleteMatch(ctx context.Context, id int64) (err error) {
	err = s.repo.DeleteMatchById(ctx, id)
	if err != nil {
		s.logger.Error("failed to delete match", zap.Error(err))
		return
	}

	return nil
}

func (s *Service) ValidateDeleteMatchId(ctx context.Context, id, issuedId int64) (err error) {
	// check issuedId and id match
	_, err = s.repo.GetMatchByIdAndIssuedId(ctx, id, issuedId)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return errors.New("matchId not found or user is not the owner of match")
	}

	return nil
}

func (s *Service) ValidateMatchIsApproved(ctx context.Context, id, issuedId int64) (err error) {
	// check issuedId and id match
	match, err := s.repo.GetMatchByIdAndIssuedId(ctx, id, issuedId)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return
	}

	if match.Status == model.MatchStatusApproved {
		return errors.New("matchId is already approved / reject")
	}

	return nil
}

func (s *Service) ApproveMatch(ctx context.Context, id int64, receiverID int64) (matchID string, err error) {
	// get match data
	data, err := s.repo.GetMatchByID(ctx, id)
	if err != nil {
		return "", err
	}

	matchCat, err := s.repo.GetCatByID(ctx, data.MatchCatId)
	if err != nil {
		return "", err
	}
	if matchCat.OwnerId != receiverID {
		return "", cerror.New(http.StatusBadRequest, "userCatId is not belong to the user")
	}

	// TODO: implement transaction
	// approve the match
	err = s.repo.UpdateMatchStatus(ctx, id, model.MatchStatusApproved)
	if err != nil {
		return "", cerror.New(http.StatusInternalServerError, "failed to update match status")
	}

	var listMatch []model.Match
	// delete the others related UserCatId MatchCatId
	listA, err := s.repo.GetMatchByUserCatIds(ctx, []int64{data.MatchCatId, data.UserCatId})
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return "", cerror.New(http.StatusInternalServerError, "failed getting match by both owner")
	}
	listMatch = append(listMatch, listA...)
	listB, err := s.repo.GetMatchByMatchCatIds(ctx, []int64{data.MatchCatId, data.UserCatId})
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return "", cerror.New(http.StatusInternalServerError, "failed getting match by both owner")
	}
	listMatch = listB

	// delete
	for _, match := range listMatch {
		// if same with the approved id, skip
		if match.ID == id {
			continue
		}

		if match.Status != model.MatchStatusWaitingForApproval {
			continue
		}
		// TODO: make this process into bulk by matchIds
		err = s.repo.DeleteMatchById(ctx, match.ID)
		if err != nil {
			s.logger.Error("failed to delete match", zap.Error(err))
		}
	}
	return
}

func (s *Service) RejectMatch(ctx context.Context, id int64) (matchID string, err error) {
	err = s.repo.UpdateMatchStatus(ctx, id, model.MatchStatusRejected)
	if err != nil {
		return matchID, cerror.New(http.StatusInternalServerError, "failed to update match status")
	}
	return
}

func (s *Service) GetMatchData(ctx context.Context, id int64) (listMatch []model.MatchData, err error) {
	var listData []model.MatchData

	rows, err := s.repo.GetAllMatchData(ctx, id)

	if err != nil {
		return nil, cerror.New(http.StatusInternalServerError, "failed getting match data")
	}

	defer rows.Close()
	for rows.Next() {
		var matchData model.MatchData
		var match model.Match
		var matchCat, userCat model.Cat
		var issuedBy model.UserResponse

		err = rows.StructScan(&match)
		if err != nil {
			return
		}

		issuedBy, _ = s.repo.GetUserById(ctx, match.IssuedID)
		matchCat, _ = s.repo.GetCatByID(ctx, match.MatchCatId)
		userCat, _ = s.repo.GetCatByID(ctx, match.UserCatId)

		matchData = model.MatchData{
			ID:             strconv.FormatInt(match.ID, 10),
			IssuedBy:       issuedBy,
			MatchCatDetail: matchCat,
			UserCatDetail:  userCat,
			Message:        match.Message,
			CreatedAt:      match.CreatedAt,
		}

		listData = append(listData, matchData)
	}

	return listData, nil
}
