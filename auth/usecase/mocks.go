package usecase

import (
	"github.com/Gidraff/task-manager-service/model"
	"github.com/stretchr/testify/mock"
)

type AuthUseCaseMock struct {
	mock.Mock
}

func (m *AuthUseCaseMock) GetAccountByEmail(email string) (res *model.Account, err error) {
	ret := m.Called(email)
	return res, ret.Error(0)
}

func (m *AuthUseCaseMock) RegisterAccount(username, email, password string) error {

	ret := m.Called(username, email, password)
	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Account) error); ok {
		r0 = rf(&model.Account{Username: username, Email: email, Password: password})
	} else {
		r0 = ret.Error(0)
	}
	return r0
}
