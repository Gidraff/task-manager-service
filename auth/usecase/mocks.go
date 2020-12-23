package usecase

import (
	"github.com/Gidraff/task-manager-service/model"
	//"github.com/Gidraff/task-manager-service/pkg/utils/helpers"
	"github.com/stretchr/testify/mock"
)

// AuthUseCaseMock type
type AuthUseCaseMock struct {
	mock.Mock
}

// GetUserByEmail mocks GetAccountByEmail
func (m *AuthUseCaseMock) GetUserByEmail(email string) (*model.User, error) {
	ret := m.Called(email)
	var r0 *model.User
	if rf, ok := ret.Get(0).(func(email string) *model.User); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(*model.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(email string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// Register  provides mock
func (m *AuthUseCaseMock) Register(username, email, password string) error {

	ret := m.Called(username, email, password)
	var r0 error
	if rf, ok := ret.Get(0).(func(*model.User) error); ok {
		r0 = rf(&model.User{Username: username, Email: email, Password: password})
	} else {
		r0 = ret.Error(0)
	}
	return r0
}
