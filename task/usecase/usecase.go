package usecase

import (
	"fmt"

	"github.com/Gidraff/task-manager-service/model"
	"github.com/Gidraff/task-manager-service/task"
)

// TaskUsecase encapsulate taskRepo
type TaskUsecase struct {
	taskRepo task.ITaskRepository
}

// NewTaskUsecase retuns a new task use case
func NewTaskUsecase(taskRepo task.ITaskRepository) *TaskUsecase {
	return &TaskUsecase{taskRepo}
}

// CreateTask saves a new task
func (taskUC *TaskUsecase) CreateTask(projectID, userID uint64, title, description, priority, status string) error {
	err := taskUC.taskRepo.Save(projectID, userID, title, description, priority, status)
	if err != nil {
		return err
	}
	return nil
}

// GetAllTasks retuns All tas
func (taskUC *TaskUsecase) GetAllTasks(projectID uint64) (*[]model.Task, error) {
	tasks, err := taskUC.taskRepo.FetchAll(projectID)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// GetTaskByID return task with id
func (taskUC *TaskUsecase) GetTaskByID(taskID, projectID uint64) (*model.Task, error) {
	task, err := taskUC.taskRepo.FetchByID(taskID, projectID)
	if err != nil {
		return nil, err
	}
	return task, nil
}

// GetTaskByTitle return task with id
func (taskUC *TaskUsecase) GetTaskByTitle(title string, projectID uint64) (*model.Task, error) {
	task, err := taskUC.taskRepo.FetchByTitle(title, projectID)
	if err != nil {
		return nil, err
	}
	return task, nil
}

// GetTaskByStatus retuns All tas
func (taskUC *TaskUsecase) GetTaskByStatus(status string, projectID uint64) (*[]model.Task, error) {
	tasks, err := taskUC.taskRepo.FetchByStatus(status, projectID)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// GetTaskByPriority retuns All tas
func (taskUC *TaskUsecase) GetTaskByPriority(priority string, projectID uint64) (*[]model.Task, error) {
	tasks, err := taskUC.taskRepo.FetchByPriority(priority, projectID)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// UpdateTask updates tasks with provded id
func (taskUC *TaskUsecase) UpdateTask(projectID, taskID uint64, title, description, priority, status string) (*model.Task, error) {
	task, err := taskUC.taskRepo.Update(taskID, projectID, title, description, priority, status)
	if err != nil {
		return nil, err
	}
	fmt.Println("===> usecase", task)
	return nil, nil
}

// DeleteTask deletes tasks using id provided
func (taskUC *TaskUsecase) DeleteTask(taskID, projectID uint64) error {
	err := taskUC.taskRepo.Delete(taskID, projectID)
	if err != nil {
		return err
	}
	return nil
}
