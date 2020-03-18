package service

import (
	"context"
	"log"

	"github.com/Gidraff/task-manager-service/cmd/taskman/model"
	"github.com/Gidraff/task-manager-service/cmd/taskman/repository"
)

// Service encapsulates user service
type Service interface {
	Register(ctx context.Context, u *model.User) error
}

type service struct {
	repo repository.UserRepository
}

// NewService creates a new user service
func NewService(repo repository.UserRepository) Service {
	return service{repo}
}

func (s service) Register(ctx context.Context, u *model.User) error {
	err := s.repo.Create(ctx, u)
	if err != nil {
		log.Printf("service %s", err)
		return err
	}

	return nil
}
