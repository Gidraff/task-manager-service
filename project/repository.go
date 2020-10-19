package project

import "github.com/Gidraff/task-manager-service/model"

// IProjectRepository define interface
type IProjectRepository interface {
	Save(userID uint64, name, description string) error
	FetchAll(userID uint64) ([]model.Project, error)
	FetchByID(projectID, userID uint64) (*model.Project, error)
	FetchByName(projectName string, userID uint64) (*model.Project, error)
	Update(projectID, userID uint64, name, description string) error
	Delete(projectID, userID uint64) error
}
