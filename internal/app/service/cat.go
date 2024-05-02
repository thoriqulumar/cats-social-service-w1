package service

import (
	"context"
	"errors"
	"strings"

	"github.com/thoriqulumar/cats-social-service-w1/internal/app/model"
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

func (s *Service) GetCat(ctx context.Context, catReq model.GetCatRequest) ([]model.Cat, error) {
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
		query += " AND has_matched = ?"
		args = append(args, *catReq.HasMatched)
	}
	if catReq.AgeInMonth != nil {
		operator, value, err := parseAgeInMonthFilter(*catReq.AgeInMonth)
		if err != nil {
			return nil, err
		}
		query += " AND age_in_month " + operator + " ?"
		args = append(args, value)
	}
	if catReq.Owned != nil {
		query += " AND owned = ?"
		args = append(args, *catReq.Owned)
	}
	if catReq.Search != nil {
		query += " AND name LIKE ?"
		args = append(args, "%"+*catReq.Search+"%")
	}
	if catReq.Limit != nil {
		query += " LIMIT ?"
		limit = *catReq.Limit
		args = append(args, limit)
	}
	if catReq.Offset != nil {
		query += " OFFSET ?"
		offset = *catReq.Offset
		args = append(args, offset)
	}

	data, err := s.repo.GetCat(ctx, query, args)
	if err != nil {
		return []model.Cat{}, err
	}

	return data, nil
}
