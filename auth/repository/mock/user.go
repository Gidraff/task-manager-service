package mock

import (
	"github.com/Gidraff/task-manager-service/model"
	"github.com/stretchr/testify/mock"
)

type UserRepoMock struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: userData
func (m *UserRepoMock) CreateUser(userData *model.User) error {
	ret := m.Called(userData)

	var r0 error
	if rf, ok := ret.Get(0).(func(user *model.User) error); ok {
		r0 = rf(userData)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}
