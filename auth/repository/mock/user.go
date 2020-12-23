package mock

import (
	"github.com/Gidraff/task-manager-service/model"
	"github.com/stretchr/testify/mock"
)

type UserRepoMock struct {
	mock.Mock
}

func (m *UserRepoMock) FetchByEmail(email string) (user *model.Account, err error) {
	ret := m.Called(email)
	return &model.Account{}, ret.Error(0)
}

// CreateUser provides a mock function with given fields: userData
func (m *UserRepoMock) Store(username, email, password string) error {
	ret := m.Called(username, email, password)
	var r0 error
	if rf, ok := ret.Get(0).(func(user *model.Account) error); ok {
		r0 = rf(&model.Account{Username: username, Email: email, Password: password})
	} else {
		r0 = ret.Error(0)
	}
	return r0
}
