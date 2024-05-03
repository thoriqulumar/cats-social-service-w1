package service

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"reflect"
	"strings"

	"github.com/thoriqulumar/cats-social-service-w1/internal/app/model"
	cerror "github.com/thoriqulumar/cats-social-service-w1/internal/pkg/error"
	"github.com/thoriqulumar/cats-social-service-w1/internal/pkg/validator"
)

func parseAgeInMonthFilter(ageFilter string) (string, string, error) {
	var operator, value string

	parts := strings.Split(ageFilter, "=")
	if len(parts) != 2 {
		return "", "", errors.New("invalid ageInMonth filter format")
	}

	operator = "="
	value = parts[1]

	if strings.HasPrefix(parts[1], "<") {
		operator = "<"
		value = parts[1][1:]
	} else if strings.HasPrefix(parts[1], ">") {
		operator = ">"
		value = parts[1][1:]
	}

	return operator, value, nil
}

func (s *Service) GetCat(ctx context.Context, catReq model.GetCatRequest, userId int64) ([]model.Cat, error) {
	limit := 5
	offset := 0

	var query string
	var args []interface{}

	if catReq.ID != nil {
		query += " AND id = ?"
		args = append(args, *catReq.ID)
	}
	if catReq.Sex != nil {
		query += " AND sex = ?"
		args = append(args, *catReq.Sex)
	}
	if catReq.Race != nil {
		query += " AND race = ?"
		args = append(args, *catReq.Race)
	}
	if catReq.HasMatched != nil {
		query += " AND hasMatched = ?"
		args = append(args, *catReq.HasMatched)
	}
	if catReq.AgeInMonth != nil {
		operator, value, err := parseAgeInMonthFilter(*catReq.AgeInMonth)
		if err != nil {
			return nil, err
		}
		query += " AND ageInMonth " + operator + " ?"
		args = append(args, value)
	}
	if catReq.Owned != nil {
		if *catReq.Owned {
			query += " AND ownerId = ?" // Get cats with ownerId equal to request's ownerId
		} else {
			query += " AND ownerId != ?" // Get cats with ownerId not equal to request's ownerId
		}
		args = append(args, userId)
	}
	if catReq.Search != nil {
		query += " AND name LIKE ?"
		args = append(args, "%"+*catReq.Search+"%")
	}

	query += " LIMIT $1 OFFSET $2"
	if catReq.Limit != nil {
		limit = *catReq.Limit
	}
	if catReq.Offset != nil {
		offset = *catReq.Offset
	}
	args = append(args, limit, offset)
	data, err := s.repo.GetCat(ctx, query, args)
	if err != nil {
		return []model.Cat{}, err
	}

	return data, nil
}

func (s *Service) PostCat(ctx context.Context, catReq model.PostCatRequest, userId int64) (model.Cat, error) {
	var args []interface{}

	// Get the type of the struct
	reqType := reflect.TypeOf(catReq)
	// Get the value of the struct
	reqValue := reflect.ValueOf(catReq)

	args = append(args, userId)
	for i := 0; i < reqType.NumField(); i++ {
		fieldValue := reqValue.Field(i).Interface()
		args = append(args, fieldValue)
	}

	data, err := s.repo.PostCat(ctx, args)
	if err != nil {
		return model.Cat{}, err
	}

	return data, nil
}

func (s *Service) PutCat(ctx context.Context, catReq model.PostCatRequest, catId int64) (sql.Result, error) {
	var args []interface{}

	// Get the type of the struct
	reqType := reflect.TypeOf(catReq)
	// Get the value of the struct
	reqValue := reflect.ValueOf(catReq)

	for i := 0; i < reqType.NumField(); i++ {
		fieldValue := reqValue.Field(i).Interface()
		args = append(args, fieldValue)
	}
	args = append(args, catId)

	data, err := s.repo.PutCat(ctx, args)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Service) ValidatePostCat(ctx context.Context, catReq model.PostCatRequest, issuerId int64) error {
	if !validator.IsString(catReq.Name) {
		return cerror.New(http.StatusBadRequest, "name doesn’t pass validation")
	}

	if !validator.IsString(catReq.Race) {
		return cerror.New(http.StatusBadRequest, "race doesn’t pass validation")
	}

	if !validator.IsString(catReq.Sex) {
		return cerror.New(http.StatusBadRequest, "sex doesn’t pass validation")
	}

	if !validator.IsNumber(catReq.AgeInMonth) {
		return cerror.New(http.StatusBadRequest, "age doesn’t pass validation")
	}

	if !validator.IsString(catReq.Description) {
		return cerror.New(http.StatusBadRequest, "description doesn’t pass validation")
	}

	if !validator.IsValidImageUrls(catReq.ImageUrls) {
		return cerror.New(http.StatusBadRequest, "imageUrls doesn’t pass validation")
	}

	return nil
}

func (s *Service) ValidatePutCat(ctx context.Context, catReq model.PostCatRequest, catId int64, issuerId int64) error {
	catData, err := s.repo.GetCatByID(ctx, catId)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return cerror.New(http.StatusNotFound, "catId is not found")
	}

	if catData.HasMatched {
		return cerror.New(http.StatusBadRequest, "sex is edited when cat is already requested to match")
	}

	_, err = s.repo.GetCatOwnerByID(ctx, catId, issuerId)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return cerror.New(http.StatusNotFound, "issuedId is not the owner of this cat")
		}
		// Handle case where no cat was found (data is zero-value Cat)
		return cerror.New(http.StatusNotFound, "issuedId is not the owner of this cat")
	}

	if !validator.IsString(catReq.Name) {
		return cerror.New(http.StatusBadRequest, "name doesn’t pass validation")
	}

	if !validator.IsString(catReq.Race) {
		return cerror.New(http.StatusBadRequest, "race doesn’t pass validation")
	}

	if !validator.IsString(catReq.Sex) {
		return cerror.New(http.StatusBadRequest, "sex doesn’t pass validation")
	}

	if !validator.IsNumber(catReq.AgeInMonth) {
		return cerror.New(http.StatusBadRequest, "age doesn’t pass validation")
	}

	if !validator.IsString(catReq.Description) {
		return cerror.New(http.StatusBadRequest, "description doesn’t pass validation")
	}

	if !validator.IsValidImageUrls(catReq.ImageUrls) {
		return cerror.New(http.StatusBadRequest, "imageUrls doesn’t pass validation")
	}

	return nil
}
