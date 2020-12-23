package task

import "github.com/Gidraff/task-manager-service/model"

// ITaskRepository interface defins repo method
type ITaskRepository interface {
	Save(projectID, userID uint64, title, description, priority, status string) error
	FetchAll(projectID uint64) (*[]model.Task, error)
	FetchByID(taskID, projectID uint64) (*model.Task, error)
	FetchByTitle(title string, projectID uint64) (*model.Task, error)
	FetchByPriority(priority string, projectID uint64) (*[]model.Task, error)
	FetchByStatus(status string, projectID uint64) (*[]model.Task, error)
	Update(taskID, projectID uint64, title, description, priority, status string) (*model.Task, error)
	Delete(taskID, projectID uint64) error
}
