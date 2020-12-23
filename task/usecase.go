package task

import "github.com/Gidraff/task-manager-service/model"

// ITaskUsecase encapsulates task methods
type ITaskUsecase interface {
	CreateTask(projectID, userID uint64, title, description, priority, status string) error
	GetAllTasks(projectID uint64) (*[]model.Task, error)
	GetTaskByID(taskID, projectID uint64) (*model.Task, error)
	GetTaskByTitle(title string, projectID uint64) (*model.Task, error)
	GetTaskByPriority(priority string, projectID uint64) (*[]model.Task, error)
	GetTaskByStatus(status string, projectID uint64) (*[]model.Task, error)
	UpdateTask(taskID, projectID uint64, title, description, priority, status string) (*model.Task, error)
	DeleteTask(taskID, projectID uint64) error
}
