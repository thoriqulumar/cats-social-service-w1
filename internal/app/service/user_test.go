package service

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/thoriqulumar/cats-social-service-w1/internal/app/config"
	"github.com/thoriqulumar/cats-social-service-w1/internal/app/model"
	"github.com/thoriqulumar/cats-social-service-w1/internal/app/service/mocks"
	"go.uber.org/zap"
	"testing"
)

func TestService_Register(t *testing.T) {
	type args struct {
		ctx  context.Context
		data model.User
	}
	tests := []struct {
		name               string
		args               args
		mock               func(t *testing.T) *Service
		wantUserWithAccess model.UserWithAccess
		wantErr            error
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				data: model.User{
					ID:       1,
					Email:    "some@email.com",
					Name:     "some",
					Password: "mypass",
				},
			},
			mock: func(t *testing.T) *Service {
				mr := mocks.NewMockrepository(gomock.NewController(t))
				mr.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(model.User{
					ID:       1,
					Email:    "some@email.com",
					Name:     "some",
					Password: "mypass",
				}, nil)
				return &Service{
					cfg:    &config.Config{},
					logger: zap.NewExample(),
					repo:   mr,
				}
			},
			wantUserWithAccess: model.UserWithAccess{
				Email: "some@email.com",
				Name:  "some",
			},
			wantErr: nil,
		},
		{
			name: "failed",
			args: args{
				ctx: context.Background(),
				data: model.User{
					ID:       1,
					Email:    "some@email.com",
					Name:     "some",
					Password: "mypass",
				},
			},
			mock: func(t *testing.T) *Service {
				mr := mocks.NewMockrepository(gomock.NewController(t))
				mr.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(model.User{}, errors.New("test-error"))
				return &Service{
					cfg:    &config.Config{},
					logger: zap.NewExample(),
					repo:   mr,
				}
			},
			wantUserWithAccess: model.UserWithAccess{},
			wantErr:            errors.New("test-error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.mock(t)
			gotUserWithAccess, err := s.Register(tt.args.ctx, tt.args.data)
			assert.EqualValues(t, tt.wantErr, err)
			assert.EqualValues(t, tt.wantUserWithAccess.Email, gotUserWithAccess.Email)
			assert.EqualValues(t, tt.wantUserWithAccess.Name, gotUserWithAccess.Name)
		})
	}
}
