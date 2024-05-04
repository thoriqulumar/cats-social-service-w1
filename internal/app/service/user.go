package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/thoriqulumar/cats-social-service-w1/internal/app/model"
	"github.com/thoriqulumar/cats-social-service-w1/internal/pkg/email"
	cerror "github.com/thoriqulumar/cats-social-service-w1/internal/pkg/error"
	"github.com/thoriqulumar/cats-social-service-w1/internal/pkg/jwt"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Register(ctx context.Context, data model.User) (userWithAccess model.UserWithAccess, err error) {
	// bcrypt password
	// Hash a password using bcrypt with the specified number of rounds
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), s.cfg.BcryptSalt)
	if err != nil {
		return
	}

	data.Password = string(hashedPassword)

	user, err := s.repo.CreateUser(ctx, data)
	fmt.Println(user)
	if err != nil {
		s.logger.Error("failed to create user", zap.Error(err))
		return
	}

	// create JWT accessToken
	user.IDStr = strconv.FormatInt(user.ID, 10)
	accessToken, err := s.generateJWT(ctx, user)
	if err != nil {
		s.logger.Error("failed generate JWT", zap.Error(err))
		return
	}
	userWithAccess = model.UserWithAccess{
		Email:       data.Email,
		Name:        data.Name,
		AccessToken: accessToken,
	}
	return
}

func (s *Service) ValidateUser(ctx context.Context, user model.User) (err error) {
	if !email.IsValidEmail(user.Email) {
		return errors.New("invalid email")
	}

	existData, err := s.repo.GetUserByEmail(ctx, user.Email)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if existData.ID > 0 {
		return errors.New("user with this email already exists")
	}

	if len(user.Password) < 5 {
		return errors.New("password is too short")
	}
	if len(user.Password) > 15 {
		return errors.New("password is too long")
	}
	//TODO add more validation from requirement docs

	return nil
}

func (s *Service) generateJWT(ctx context.Context, user model.User) (string, error) {
	return jwt.Generate(s.cfg.JWTSecret, user)
}

func (s *Service) Login(ctx context.Context, user model.LoginRequest) (data model.UserWithAccess, err error) {
	existData, err := s.repo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		s.logger.Error("failed GetUserByEmail", zap.Error(err))
		// TODO: handle between entry error n db error, ex: userNotFound
		return data, cerror.New(http.StatusInternalServerError, err.Error())
	}

	// compare password
	err = bcrypt.CompareHashAndPassword([]byte(existData.Password), []byte(user.Password))
	if err != nil {
		s.logger.Error("failed bcrypt.GenerateFromPassword", zap.Error(err))
		return data, cerror.New(http.StatusBadRequest, err.Error())
	}

	// create JWT accessToken
	existData.IDStr = strconv.FormatInt(existData.ID, 10)
	accessToken, err := s.generateJWT(ctx, existData)
	if err != nil {
		s.logger.Error("failed generate JWT", zap.Error(err))
		return
	}

	data = model.UserWithAccess{
		Email:       existData.Email,
		Name:        existData.Name,
		AccessToken: accessToken,
	}
	return data, nil
}
