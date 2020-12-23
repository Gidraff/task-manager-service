package postgres

import (
	"fmt"
	"github.com/Gidraff/task-manager-service/model"
	//"github.com/Gidraff/task-manager-service/pkg/utils/helpers"
	"github.com/Gidraff/task-manager-service/project"
	//"github.com/lib/pq"
	"gorm.io/gorm"
	"time"
)

// QueryWthIDAndUserIDCondition common condition
const (
	Tasks                        = "Tasks"
	QueryWthIDAndUserIDCondition = "id = ? AND user_id = ?"
)

type projectRepo struct {
	Conn *gorm.DB
}

// NewProjectRepo returns a new UserRepository interface
func NewProjectRepo(db *gorm.DB) project.IProjectRepository {
	return &projectRepo{Conn: db}
}

// Save inserts a new record
func (pr *projectRepo) Save(userID uint64, name, description string) error {
	//updatedAt := helpers.NullTime{}
	project := &model.Project{
		Name:        name,
		Description: description,
		UserID:      userID,
		CreatedAt:   time.Now(),
	}
	result := pr.Conn.Create(&project)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// FetchAll returns all
func (pr *projectRepo) FetchAll(userID uint64) ([]model.Project, error) {
	var projects []model.Project
	err := pr.Conn.Preload(Tasks).
		Where("user_id = ?", userID).
		Find(&projects).Error
	if err != nil {
		return nil, err
	}
	return projects, nil
}

// FetchByID return a record matching the id
func (pr *projectRepo) FetchByID(projectID, userID uint64) (*model.Project, error) {
	project := &model.Project{}
	err := pr.Conn.Preload(Tasks).
		Where(QueryWthIDAndUserIDCondition, projectID, userID).
		First(&project).Error
	if err != nil {
		return &model.Project{}, err
	}
	return project, nil
}

// FetchByName return a record matching the name
func (pr *projectRepo) FetchByName(projectName string, userID uint64) (*model.Project, error) {
	project := &model.Project{}
	result := pr.Conn.Preload(Tasks).
		Where("name = ? and user_id = ?", projectName, userID).First(&project)
	if result.Error != nil {
		return &model.Project{}, result.Error
	}
	return project, nil
}

// Update updates a record matching the id
func (pr *projectRepo) Update(projectID, userID uint64, name, description string) error {
	var err error
	project := &model.Project{}
	err = pr.Conn.Debug().Model(project).
		Where(QueryWthIDAndUserIDCondition, projectID, userID).
		Updates(&model.Project{
			Name: name, Description: description}).Error
	if err != nil {
		return err
	}
	return nil
}

// Delete delete a record matching the id
func (pr *projectRepo) Delete(projectID, userID uint64) error {
	project := &model.Project{ID: projectID, UserID: userID}
	result := pr.Conn.Select(Tasks).
		Delete(&project)
	fmt.Println(userID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
