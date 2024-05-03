package service

import (
	"context"

	"errors"
	"fmt"
	"net/url"

	"go.uber.org/zap"

	"github.com/thoriqulumar/cats-social-service-w1/internal/app/model"
)

func (s *Service) RegisterCat(ctx context.Context, data model.Cat) (model.Cat, error) {

	cat, err := s.repo.CreateCat(ctx, data)
	if err != nil {
		s.logger.Error("failed to create user", zap.Error(err))
		return model.Cat{}, err
	}

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
