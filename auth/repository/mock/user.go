package mock

import (
	"github.com/Gidraff/task-manager-service/model"
	"github.com/stretchr/testify/mock"
)

type UserRepoMock struct {
	mock.Mock
}

func (m *UserRepoMock) FetchByEmail(email string) (user *model.User, err error) {
	ret := m.Called(email)
	return &model.User{}, ret.Error(0)
}

// CreateUser provides a mock function with given fields: userData
func (m *UserRepoMock) Store(userData *model.User) error {
	ret := m.Called(userData)
	var r0 error
	if rf, ok := ret.Get(0).(func(user *model.User) error); ok {
		r0 = rf(userData)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}
