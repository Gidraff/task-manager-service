package project

import "github.com/Gidraff/task-manager-service/model"

// IProjectUsecase define type methods
type IProjectUsecase interface {
	CreateNewProject(userID uint64, name, description string) error
	FetchAllProjects(userID uint64) ([]model.Project, error)
	FetchProjectByID(projectID, userID uint64) (*model.Project, error)
	FetchProjectByName(projectName string, userID uint64) (*model.Project, error)
	UpdateProject(projectID, userID uint64, name, description string) error
	DeleteProject(projectID, userID uint64) error
}
