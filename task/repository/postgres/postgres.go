package postgres

import (
	"fmt"

	"github.com/Gidraff/task-manager-service/model"
	"github.com/Gidraff/task-manager-service/task"
	"gorm.io/gorm"
)

// QueryWithIDAndProjectIDCondition query condtion
const QueryWithIDAndProjectIDCondition = "id = ? and project_id = ?"

type taskRepo struct {
	Conn *gorm.DB
}

// NewTaskRepo returns type Task
func NewTaskRepo(conn *gorm.DB) task.ITaskRepository {
	return &taskRepo{Conn: conn}
}

// SAVE stores task to db
func (tr *taskRepo) Save(projectID, userID uint64, title, description, priority, status string) error {
	var err error
	task := &model.Task{
		ProjectID:   projectID,
		UserID:      userID,
		Title:       title,
		Description: description,
		Priority:    priority,
		Status:      status,
	}

	result := tr.Conn.Create(&task)
	err = result.Error
	if err != nil {
		return err
	}
	return nil
}

// FetchAll returns all tasks
func (tr *taskRepo) FetchAll(projectID uint64) (*[]model.Task, error) {

	var err error
	tasks := &[]model.Task{}
	err = tr.Conn.Debug().Model(&model.Task{}).
		Where("project_id = ?", projectID).
		Limit(100).
		Find(&tasks).Error
	if err != nil {
		return &[]model.Task{}, err
	}
	return tasks, nil
}

// FetchByID return tasks with id provided
func (tr *taskRepo) FetchByID(taskID, projectID uint64) (*model.Task, error) {
	task := &model.Task{}
	result := tr.Conn.Where(QueryWithIDAndProjectIDCondition, taskID, projectID).First(&task)
	if result.Error != nil {
		return nil, result.Error
	}
	return task, nil
}

// FetchByID return tasks with id provided
func (tr *taskRepo) FetchByPriority(priority string, projectID uint64) (*[]model.Task, error) {
	var err error
	tasks := &[]model.Task{}
	err = tr.Conn.Debug().Model(&model.Task{}).
		Where("priority = ? and project_id = ?", priority, projectID).
		Limit(100).
		Find(&tasks).Error
	if err != nil {
		return &[]model.Task{}, err
	}
	return tasks, nil
}

// FetchByID return tasks with id provided
func (tr *taskRepo) FetchByStatus(status string, projectID uint64) (*[]model.Task, error) {
	var err error
	tasks := &[]model.Task{}
	err = tr.Conn.Debug().Model(&model.Task{}).
		Where("status = ? and project_id = ?", status, projectID).
		Limit(100).
		Find(&tasks).Error
	if err != nil {
		return &[]model.Task{}, err
	}
	return tasks, nil
}

// FetchByID return tasks with id provided
func (tr *taskRepo) FetchByTitle(title string, projectID uint64) (*model.Task, error) {
	task := &model.Task{}
	result := tr.Conn.Where("title = ? and project_id = ?", title, projectID).First(&task)
	if result.Error != nil {
		return nil, result.Error
	}
	return task, nil
}

// UPDATE updates task
func (tr *taskRepo) Update(taskID, projectID uint64, title, description, priority, status string) (*model.Task, error) {
	var err error
	task := &model.Task{}
	err = tr.Conn.Debug().Model(task).
		Where(QueryWithIDAndProjectIDCondition, taskID, projectID).
		Updates(
			&model.Task{
				Title:       title,
				Description: description,
				Priority:    priority,
				Status:      status}).Error
	if err != nil {
		return &model.Task{}, err
	}
	return task, nil

}

// DELETE deletes task with id provided
func (tr *taskRepo) Delete(taskID, projectID uint64) error {
	result := tr.Conn.Debug().Model(&model.Task{}).
		Where(QueryWithIDAndProjectIDCondition, taskID, projectID).
		Take(&model.Task{}).
		Delete(&model.Task{})
	if result.Error != nil {
		return result.Error
	}
	fmt.Println(result.RowsAffected)
	return nil
}
