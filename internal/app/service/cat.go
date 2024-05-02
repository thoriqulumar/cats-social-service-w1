package service

import (
	"context"
	"fmt"

	"github.com/thoriqulumar/cats-social-service-w1/internal/app/model"
)

func (s *Service) GetCat(ctx context.Context, catReq model.GetCatRequest) (model.Cat, error) {
	fmt.Println("catReq", catReq)
	// // Extract filter from request context
	// req, ok := ctx.Value("request").(*http.Request)
	// if !ok {
	// 	return model.Cat{}, errors.New("Request context not found")
	// }

	// query := req.URL.Query()

	// filter := ""

	// data, err := s.repo.GetCat(ctx, filter)
	// if err != nil {
	// 	return model.Cat{}, err
	// }

	// return data, nil
	return model.Cat{}, nil
}
