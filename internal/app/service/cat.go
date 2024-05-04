package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/thoriqulumar/cats-social-service-w1/internal/pkg/converter"
	"github.com/thoriqulumar/cats-social-service-w1/internal/pkg/validator"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/thoriqulumar/cats-social-service-w1/internal/app/model"
	cerror "github.com/thoriqulumar/cats-social-service-w1/internal/pkg/error"
)

func (s *Service) RegisterCat(ctx context.Context, data model.Cat, userId int64) (model.Cat, error) {

	data.OwnerId = userId
	cat, err := s.repo.CreateCat(ctx, data)
	if err != nil {
		s.logger.Error("failed to create user", zap.Error(err))
		return model.Cat{}, err
	}
	cat.IDStr = strconv.FormatInt(cat.ID, 10)
	cat.CreatedAt = time.Now().UTC().Format(time.RFC3339)
	return cat, nil

}
func (s *Service) ValidateCat(ctx context.Context, cat model.Cat) (err error) {

	// Validate name
	if len(cat.Name) < 1 || len(cat.Name) > 30 {
		return errors.New("name must be between 1 and 30 characters long")
	}

	// Validate race
	validRaces := []string{"Persian", "Maine Coon", "Siamese", "Ragdoll", "Bengal", "Sphynx", "British Shorthair", "Abyssinian", "Scottish Fold", "Birman"}
	isValidRace := false
	for _, race := range validRaces {
		if cat.Race == race {
			isValidRace = true
			break
		}
	}
	if !isValidRace {
		return errors.New("race is invalid or not specified")
	}

	// Validate sex
	if cat.Sex != "male" && cat.Sex != "female" {
		return errors.New("sex must be either 'male' or 'female'")
	}

	// Validate ageInMonth
	if cat.AgeInMonth < 1 || cat.AgeInMonth > 120082 {
		return fmt.Errorf("ageInMonth must be between 1 and 120082, got %d", cat.AgeInMonth)
	}

	// Validate description
	if len(cat.Description) < 1 || len(cat.Description) > 200 {
		return errors.New("description must be between 1 and 200 characters long")
	}

	// Validate imageUrls
	if len(cat.ImagesUrls) == 0 {
		return errors.New("at least one imageUrl is required")
	}
	for _, urlStr := range cat.ImagesUrls {
		if urlStr == "" {
			return errors.New("imageUrls cannot contain empty strings")
		}
		if _, err := url.ParseRequestURI(urlStr); err != nil {
			return errors.New("each imageUrl must be a valid URL")
		}
	}

	//TODO add more validation from requirement docs

	return nil
}

func parseAgeInMonthFilter(ageFilter string) (string, string, error) {
	var operator, value string

	if strings.Contains(ageFilter, ">") {
		operator = ">"
		value = strings.TrimPrefix(ageFilter, ">")
	}
	if strings.Contains(ageFilter, "<") {
		operator = "<"
		value = strings.TrimPrefix(ageFilter, "<")
	}
	if value == "" {
		operator = "="
		value = strings.TrimPrefix(ageFilter, "=")
	}

	return operator, value, nil
}

func (s *Service) GetCat(ctx context.Context, catReq model.GetCatRequest, userId int64) ([]model.Cat, error) {
	var query string
	var args []interface{}

	if catReq.ID != "" {
		query += " AND id = ?"
		// parsing to int64
		id, _ := strconv.ParseInt(catReq.ID, 10, 64)
		args = append(args, id)
	}
	if catReq.Sex != "" {
		query += " AND sex = ?"
		args = append(args, catReq.Sex)
	}
	if catReq.Race != "" {
		query += " AND race = ?"
		args = append(args, catReq.Race)
	}
	if catReq.HasMatched != nil {
		query += ` AND "hasMatched" = ?`
		args = append(args, *catReq.HasMatched)
	}
	if catReq.AgeInMonth != "" {
		operator, value, err := parseAgeInMonthFilter(catReq.AgeInMonth)
		if err != nil {
			return nil, err
		}
		query += ` AND "ageInMonth" ` + operator + " ?"
		args = append(args, value)
	}
	if catReq.Owned != nil {
		if *catReq.Owned {
			query += ` AND "ownerId" = ?` // Get cats with ownerId equal to request's ownerId
		} else {
			query += ` AND "ownerId" != ?` // Get cats with ownerId not equal to request's ownerId
		}
		args = append(args, userId)
	}
	if catReq.Search != "" {
		query += " AND name LIKE ?"
		args = append(args, "%"+catReq.Search+"%")
	}

	query += fmt.Sprintf(` LIMIT %d OFFSET %d`, catReq.Limit, catReq.Offset)
	data, err := s.repo.GetCat(ctx, query, args)
	if err != nil {
		return []model.Cat{}, err
	}

	return data, nil
}

func (s *Service) PutCat(ctx context.Context, catReq model.PostCatRequest, catId int64) (sql.Result, error) {
	var args []interface{}

	inputVal := reflect.ValueOf(catReq)
	for i := 0; i < inputVal.NumField()-1; i++ {
		args = append(args, inputVal.Field(i).Interface())
	}
	args = append(args, converter.ConvertStrArrToPgArr(catReq.ImageUrls))
	args = append(args, catId)

	data, err := s.repo.PutCat(ctx, args)
	if err != nil {
		return nil, err
	}

	return data, nil
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

func (s *Service) DeleteCat(ctx context.Context, id int64) (err error) {
	err = s.repo.DeleteCatById(ctx, id)
	fmt.Println(err)
	if err != nil {
		s.logger.Error("failed to delete cat", zap.Error(err))
		return
	}

	return nil
}

func (s *Service) ValidateDeleteCat(ctx context.Context, id, issuedId int64) (err error) {
	_, err = s.repo.GetCatOwnerByID(ctx, id, issuedId)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return cerror.New(http.StatusBadRequest, "catId not found or user is not the owner of cat")
	}

	return nil
}
