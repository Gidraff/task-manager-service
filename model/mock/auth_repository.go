package mock

import (
	"github.com/Gidraff/task-manager-service/model"
	"github.com/stretchr/testify/mock"
)

// UserRepoMock encapsulates mock
type UserRepoMock struct {
	mock.Mock
}

// NewUserRepoMock returns a new userRepo mock
func NewUserRepoMock() *UserRepoMock {
	return &UserRepoMock{}
}

// FetchByEmail provides a mock function
func (m *UserRepoMock) FetchByEmail(email string) (res model.User, err error) {
	ret := m.Called(email)

	if rf, ok := ret.Get(0).(func(email string) model.User); ok {
		res = rf(email)
	} else {
		if ret.Get(0) != nil {
			res = ret.Get(0).(model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(email string) error); ok {
		err = rf(email)
	} else {
		err = ret.Error(1)
	}
	return res, err
}

// Store provides a mock function with given fields: userData
func (m *UserRepoMock) Store(username, email, password string) error {
	ret := m.Called(username, email, password)
	var r0 error
	if rf, ok := ret.Get(0).(func(user *model.User) error); ok {
		r0 = rf(&model.User{Username: username, Email: email, Password: password})
	} else {
		r0 = ret.Error(0)
	}
	return r0
}
