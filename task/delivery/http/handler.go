package http

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Gidraff/task-manager-service/model"
	"github.com/Gidraff/task-manager-service/pkg/utils/helpers"
	"github.com/Gidraff/task-manager-service/task"
	//"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

const (
	// ServerError server error message
	ServerError = "Something went wrong while processing your request"
	// ResourceNotFound message
	ResourceNotFound = "Resource not found"
)

// TaskInput type
type TaskInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Status      string `json:"status"`
}

// TaskHandler encapsulate Task usecase
type TaskHandler struct {
	taskUC task.ITaskUsecase
}

// NewTaskHandler return Type TaskHandler
func NewTaskHandler(taskUC task.ITaskUsecase) *TaskHandler {
	return &TaskHandler{taskUC}
}

// CreateTask handles POST request
func (th *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pID, err := strconv.Atoi(vars["projectID"])
	if err != nil {
		helpers.Response(http.StatusInternalServerError, ServerError, w)
		return
	}
	var task *model.Task
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	if err := decoder.Decode(&task); err != nil {
		msg := "Request body contains badly-formatted JSON (at position %d)"
		helpers.Response(http.StatusBadRequest, msg, w)
		return
	}
	userID, err := helpers.ExtractTokenMetadata(r)
	if err != nil {
		helpers.Response(http.StatusInternalServerError, ServerError, w)
		return
	}

	isTitleValid, err := helpers.ValidateTaskTitle(task.Title)
	if !isTitleValid {
		helpers.Response(http.StatusBadRequest, err.Error(), w)
		return
	}
	duplicateTask, _ := th.taskUC.GetTaskByTitle(task.Title, uint64(pID))
	if duplicateTask != nil {
		helpers.Response(http.StatusConflict, "Resource already exist", w)
		return
	}
	err = th.taskUC.CreateTask(uint64(pID), userID, task.Title, task.Description, task.Priority, task.Status)
	if err != nil {
		helpers.Response(http.StatusInternalServerError, ServerError, w)
		return
	}
	helpers.Response(201, "Resource successfully created", w)
	return
}

// GetAllTasks return Type Tasks
func (th *TaskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pID, err := strconv.Atoi(vars["projectID"])
	if err != nil {
		helpers.Response(http.StatusInternalServerError, ServerError, w)
		return
	}
	tasks, err := th.taskUC.GetAllTasks(uint64(pID))
	if err != nil {
		helpers.Response(http.StatusInternalServerError, ServerError, w)
		return
	}
	helpers.PayloadResponse(tasks, w)
	return
}

// GetTask return task matching id
func (th *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tID, err := strconv.Atoi(vars["taskID"])
	if err != nil {
		helpers.Response(http.StatusInternalServerError, ServerError, w)
		return
	}
	pID, err := strconv.Atoi(vars["projectID"])
	if err != nil {
		helpers.Response(http.StatusInternalServerError, ServerError, w)
		return
	}
	task, err := th.taskUC.GetTaskByID(uint64(tID), uint64(pID))
	if err != nil {
		helpers.Response(http.StatusNotFound, ResourceNotFound, w)
		return
	}
	helpers.PayloadResponse(task, w)
	return
}

// GetTaskByTitle handle get by title
func (th *TaskHandler) GetTaskByTitle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pID, err := strconv.Atoi(vars["projectID"])
	if err != nil {
		helpers.Response(http.StatusInternalServerError, ServerError, w)
		return
	}
	title := r.URL.Query().Get("title")
	if title == "" {
		helpers.Response(http.StatusBadRequest, "Invalid query parameter", w)
		return
	}
	task, err := th.taskUC.GetTaskByTitle(title, uint64(pID))
	if err != nil {
		helpers.Response(http.StatusBadRequest, ResourceNotFound, w)
		return
	}
	helpers.PayloadResponse(task, w)
	return
}

// FilterTaskByStatus filter tasks by status
func (th *TaskHandler) FilterTaskByStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectID, err := strconv.Atoi(vars["projectID"])
	if err != nil {
		helpers.Response(http.StatusInternalServerError, ServerError, w)
		return
	}
	filter := r.URL.Query().Get("status")
	if filter == "" {
		helpers.Response(http.StatusBadRequest, "Invalid query parameter", w)
		return
	}
	tasks, err := th.taskUC.GetTaskByStatus(filter, uint64(projectID))
	if err != nil {
		helpers.Response(http.StatusBadRequest, ResourceNotFound, w)
		return
	}
	helpers.PayloadResponse(tasks, w)
	return
}

// FilterTaskByPriority filter tasks by priority
func (th *TaskHandler) FilterTaskByPriority(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectID, err := strconv.Atoi(vars["projectID"])
	if err != nil {
		log.Println(", err", err)
	}
	filter := r.URL.Query().Get("priority")
	if filter == "" {
		helpers.Response(http.StatusBadRequest, "Invalid request parameter", w)
		return
	}
	tasks, err := th.taskUC.GetTaskByPriority(filter, uint64(projectID))
	if err != nil {
		helpers.Response(http.StatusBadRequest, ResourceNotFound, w)
		return
	}
	helpers.PayloadResponse(tasks, w)
	return
}

// UpdateTask handle updates request
func (th *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tID, err := strconv.Atoi(vars["taskID"])
	if err != nil {
		helpers.Response(http.StatusInternalServerError, ServerError, w)
		return
	}
	pID, err := strconv.Atoi(vars["projectID"])
	if err != nil {
		helpers.Response(http.StatusInternalServerError, ServerError, w)
		return
	}

	var task *model.Task
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	if err := decoder.Decode(&task); err != nil {
		msg := "Request body contains badly-formatted JSON (at position %d)"
		helpers.Response(http.StatusBadRequest, msg, w)
		return
	}
	isTitleValid, err := helpers.ValidateTaskTitle(task.Title)
	if !isTitleValid {
		helpers.Response(http.StatusBadRequest, err.Error(), w)
		return
	}
	task, err = th.taskUC.UpdateTask(uint64(pID), uint64(tID), task.Title, task.Description, task.Priority, task.Status)
	if err != nil {
		helpers.Response(http.StatusInternalServerError, ServerError, w)
		return
	}
	helpers.Response(201, "Resource successfully updated", w)
	return
}

// DeleteTask handle delete request
func (th *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	var err error
	var tID int
	var pID int

	vars := mux.Vars(r)
	tID, err = strconv.Atoi(vars["taskID"])
	if err != nil {
		helpers.Response(http.StatusInternalServerError, ServerError, w)
		return
	}
	pID, err = strconv.Atoi(vars["projectID"])
	if err != nil {
		helpers.Response(http.StatusInternalServerError, ServerError, w)
		return
	}
	err = th.taskUC.DeleteTask(uint64(tID), uint64(pID))
	if err != nil {
		helpers.Response(http.StatusInternalServerError, ResourceNotFound, w)
		return
	}
	helpers.Response(http.StatusOK, "Resource successfully deleted", w)
	return
}
