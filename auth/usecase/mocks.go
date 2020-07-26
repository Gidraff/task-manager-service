package usecase

import (
	"github.com/Gidraff/task-manager-service/model"
	"github.com/stretchr/testify/mock"
)

type AuthUseCaseMock struct {
	mock.Mock
}

func (m *AuthUseCaseMock) GetUserByEmail(email string) (res *model.User, err error) {
	ret := m.Called(email)
	return res, ret.Error(0)
}

func (m *AuthUseCaseMock) SignUp(user *model.User) error {
	ret := m.Called(user)
	var r0 error
	if rf, ok := ret.Get(0).(func(*model.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}
