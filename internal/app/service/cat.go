package service

import (
	"context"
	"fmt"

	"github.com/thoriqulumar/cats-social-service-w1/internal/app/model"
)

func (s *Service) GetCat(ctx context.Context, catReq model.GetCatRequest) ([]model.Cat, error) {
	fmt.Println("catReq", catReq)
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
		query += " AND ageInMonth = ?"
		args = append(args, *catReq.AgeInMonth)
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
