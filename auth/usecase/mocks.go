package usecase

import (
	"github.com/Gidraff/task-manager-service/model"
	"github.com/stretchr/testify/mock"
)

type AuthUseCaseMock struct {
	mock.Mock
}

func (m *AuthUseCaseMock) SignUp(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}
