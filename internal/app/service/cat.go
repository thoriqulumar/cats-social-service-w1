package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/thoriqulumar/cats-social-service-w1/internal/app/model"
)

func (s *Service) GetCat(ctx context.Context, catReq model.GetCatRequest) (model.Cat, error) {
	fmt.Println("catReq", catReq)
	limit := 5
	offset := 0
	// // Extract filter from request context
	// req, ok := ctx.Value("request").(*http.Request)
	// if !ok {
	// 	return model.Cat{}, errors.New("Request context not found")
	// }

	// query := req.URL.Query()

	filter := ""

	if catReq.ID != nil {
		filter += "id = " + *catReq.ID
	}
	if catReq.Sex != nil {
		filter += "sex = " + *catReq.Sex
	}
	if catReq.Race != nil {
		filter += "race = " + *catReq.Race
	}
	if catReq.HasMatched != nil {
		filter += "hasMatched = " + strconv.FormatBool(*catReq.HasMatched)
	}
	if catReq.AgeInMonth != nil {
		filter += "ageInMonth = " + *catReq.AgeInMonth
	}

	if catReq.Limit != nil {
		limit = *catReq.Limit
	}
	if catReq.Offset != nil {
		offset = *catReq.Offset
	}

	data, err := s.repo.GetCat(ctx, limit, offset)
	if err != nil {
		return model.Cat{}, err
	}

	return data, nil
}
