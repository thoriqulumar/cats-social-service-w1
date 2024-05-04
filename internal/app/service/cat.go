package service

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/thoriqulumar/cats-social-service-w1/internal/app/model"
	"go.uber.org/zap"
)

func (s *Service) RegisterCat(ctx context.Context, data model.Cat, userId int64) (model.Cat, error) {

	data.OwnerId = userId
	cat, err := s.repo.CreateCat(ctx, data)
	fmt.Println(cat)
	if err != nil {
		s.logger.Error("failed to create user", zap.Error(err))
		return model.Cat{}, err
	}
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
	if len(cat.ImagesUrl) == 0 {
		return errors.New("at least one imageUrl is required")
	}
	for _, urlStr := range cat.ImagesUrl {
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
