package usecase

import (
	"log"

	"github.com/Gidraff/task-manager-service/model"
	"github.com/Gidraff/task-manager-service/project"
)

// ProjectUsecase type
type ProjectUsecase struct {
	repository project.IProjectRepository
}

// NewProjectUsecase retuns a new project use case
func NewProjectUsecase(projectRepo project.IProjectRepository) *ProjectUsecase {
	return &ProjectUsecase{projectRepo}
}

// CreateNewProject create a new project
func (pu *ProjectUsecase) CreateNewProject(userID uint64, name, description string) error {
	err := pu.repository.Save(userID, name, description)
	if err != nil {
		return err
	}
	return nil
}

// FetchAllProjects return all projects
func (pu *ProjectUsecase) FetchAllProjects(userID uint64) ([]model.Project, error) {
	projects, err := pu.repository.FetchAll(userID)
	if err != nil {
		return nil, err
	}
	log.Println("usecase project", projects)
	return projects, nil
}

// FetchProjectByID return project by id
func (pu *ProjectUsecase) FetchProjectByID(projectID, userID uint64) (*model.Project, error) {
	project, err := pu.repository.FetchByID(projectID, userID)
	if err != nil {
		log.Printf("project usecase: an error occured - %s", err)
		return nil, err
	}
	return project, nil
}

// FetchProjectByName return project by name
func (pu *ProjectUsecase) FetchProjectByName(projectName string, userID uint64) (*model.Project, error) {
	project, err := pu.repository.FetchByName(projectName, userID)
	if err != nil {
		log.Printf("Project: failed to fetch --> %s", err)
		return nil, err
	}
	return project, nil
}

// UpdateProject updates project matching id
func (pu *ProjectUsecase) UpdateProject(projectID, userID uint64, name, description string) error {
	log.Println("Project: updating...")
	err := pu.repository.Update(projectID, userID, name, description)
	if err != nil {
		log.Printf("Project: failed to update --> %s", err)
		return err
	}
	return nil
}

// DeleteProject delete project with id provided
func (pu *ProjectUsecase) DeleteProject(projectID, userID uint64) error {
	log.Printf("usecase: deleting project with id {%d}", projectID)
	err := pu.repository.Delete(projectID, userID)
	if err != nil {
		log.Printf("Project: failed to delete --> %s", err)
		return err
	}
	log.Printf("Project: successfully deleted")
	return nil
}
