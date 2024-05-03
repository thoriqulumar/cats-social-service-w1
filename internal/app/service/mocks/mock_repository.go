// Code generated by MockGen. DO NOT EDIT.
// Source: init.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/thoriqulumar/cats-social-service-w1/internal/app/model"
)

// Mockrepository is a mock of repository interface.
type Mockrepository struct {
	ctrl     *gomock.Controller
	recorder *MockrepositoryMockRecorder
}

// MockrepositoryMockRecorder is the mock recorder for Mockrepository.
type MockrepositoryMockRecorder struct {
	mock *Mockrepository
}

// NewMockrepository creates a new mock instance.
func NewMockrepository(ctrl *gomock.Controller) *Mockrepository {
	mock := &Mockrepository{ctrl: ctrl}
	mock.recorder = &MockrepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockrepository) EXPECT() *MockrepositoryMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *Mockrepository) CreateUser(ctx context.Context, data model.User) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, data)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockrepositoryMockRecorder) CreateUser(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*Mockrepository)(nil).CreateUser), ctx, data)
}

// DeleteMatchById mocks base method.
func (m *Mockrepository) DeleteMatchById(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMatchById", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMatchById indicates an expected call of DeleteMatchById.
func (mr *MockrepositoryMockRecorder) DeleteMatchById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMatchById", reflect.TypeOf((*Mockrepository)(nil).DeleteMatchById), ctx, id)
}

// GetCatByID mocks base method.
func (m *Mockrepository) GetCatByID(ctx context.Context, id int64) (model.Cat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCatByID", ctx, id)
	ret0, _ := ret[0].(model.Cat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCatByID indicates an expected call of GetCatByID.
func (mr *MockrepositoryMockRecorder) GetCatByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCatByID", reflect.TypeOf((*Mockrepository)(nil).GetCatByID), ctx, id)
}

// GetCatOwnerByID mocks base method.
func (m *Mockrepository) GetCatOwnerByID(ctx context.Context, catId, ownerId int64) (model.Cat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCatOwnerByID", ctx, catId, ownerId)
	ret0, _ := ret[0].(model.Cat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCatOwnerByID indicates an expected call of GetCatOwnerByID.
func (mr *MockrepositoryMockRecorder) GetCatOwnerByID(ctx, catId, ownerId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCatOwnerByID", reflect.TypeOf((*Mockrepository)(nil).GetCatOwnerByID), ctx, catId, ownerId)
}

// GetMatchByID mocks base method.
func (m *Mockrepository) GetMatchByID(ctx context.Context, id int64) (model.Match, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMatchByID", ctx, id)
	ret0, _ := ret[0].(model.Match)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMatchByID indicates an expected call of GetMatchByID.
func (mr *MockrepositoryMockRecorder) GetMatchByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMatchByID", reflect.TypeOf((*Mockrepository)(nil).GetMatchByID), ctx, id)
}

// GetMatchByIdAndIssuedId mocks base method.
func (m *Mockrepository) GetMatchByIdAndIssuedId(ctx context.Context, id, issuedId int64) (model.Match, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMatchByIdAndIssuedId", ctx, id, issuedId)
	ret0, _ := ret[0].(model.Match)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMatchByIdAndIssuedId indicates an expected call of GetMatchByIdAndIssuedId.
func (mr *MockrepositoryMockRecorder) GetMatchByIdAndIssuedId(ctx, id, issuedId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMatchByIdAndIssuedId", reflect.TypeOf((*Mockrepository)(nil).GetMatchByIdAndIssuedId), ctx, id, issuedId)
}

// GetMatchByMatchCatIds mocks base method.
func (m *Mockrepository) GetMatchByMatchCatIds(ctx context.Context, matchCatIDs []int64) ([]model.Match, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMatchByMatchCatIds", ctx, matchCatIDs)
	ret0, _ := ret[0].([]model.Match)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMatchByMatchCatIds indicates an expected call of GetMatchByMatchCatIds.
func (mr *MockrepositoryMockRecorder) GetMatchByMatchCatIds(ctx, matchCatIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMatchByMatchCatIds", reflect.TypeOf((*Mockrepository)(nil).GetMatchByMatchCatIds), ctx, matchCatIDs)
}

// GetMatchByUserCatIds mocks base method.
func (m *Mockrepository) GetMatchByUserCatIds(ctx context.Context, userCatIDs []int64) ([]model.Match, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMatchByUserCatIds", ctx, userCatIDs)
	ret0, _ := ret[0].([]model.Match)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMatchByUserCatIds indicates an expected call of GetMatchByUserCatIds.
func (mr *MockrepositoryMockRecorder) GetMatchByUserCatIds(ctx, userCatIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMatchByUserCatIds", reflect.TypeOf((*Mockrepository)(nil).GetMatchByUserCatIds), ctx, userCatIDs)
}

// GetUserByEmail mocks base method.
func (m *Mockrepository) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", ctx, email)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockrepositoryMockRecorder) GetUserByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*Mockrepository)(nil).GetUserByEmail), ctx, email)
}

// MatchCat mocks base method.
func (m *Mockrepository) MatchCat(ctx context.Context, data model.MatchRequest, issuedId, receiverID int64) (model.Match, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MatchCat", ctx, data, issuedId, receiverID)
	ret0, _ := ret[0].(model.Match)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MatchCat indicates an expected call of MatchCat.
func (mr *MockrepositoryMockRecorder) MatchCat(ctx, data, issuedId, receiverID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MatchCat", reflect.TypeOf((*Mockrepository)(nil).MatchCat), ctx, data, issuedId, receiverID)
}

// UpdateMatchStatus mocks base method.
func (m *Mockrepository) UpdateMatchStatus(ctx context.Context, id int64, status model.MatchStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMatchStatus", ctx, id, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMatchStatus indicates an expected call of UpdateMatchStatus.
func (mr *MockrepositoryMockRecorder) UpdateMatchStatus(ctx, id, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMatchStatus", reflect.TypeOf((*Mockrepository)(nil).UpdateMatchStatus), ctx, id, status)
}