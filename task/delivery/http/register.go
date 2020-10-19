package http

import (
	"github.com/Gidraff/task-manager-service/task"
	"github.com/gorilla/mux"
)

// RegisterHandler register endpoint
func RegisterHandler(router *mux.Router, uc task.ITaskUsecase) {
	const commonPath = "/api/v1/project/{projectID}/task/{taskID}"
	h := NewTaskHandler(uc)
	router.HandleFunc("/api/v1/project/{projectID}/task/", h.CreateTask).Methods("POST")
	router.HandleFunc("/api/v1/project/{projectID}/task/", h.GetAllTasks).Methods("GET")
	router.HandleFunc(commonPath, h.GetTask).Methods("GET")
	router.HandleFunc("/api/v1/project/{projectID}/task", h.GetTaskByTitle).
		Queries("title", "{title}")
	router.HandleFunc("/api/v1/project/{projectID}/task", h.FilterTaskByStatus).
		Methods("GET").
		Queries("status", "{status}")
	router.HandleFunc("/api/v1/project/{projectID}/task", h.FilterTaskByPriority).
		Methods("GET").
		Queries("priority", "{priority}")
	router.HandleFunc(commonPath, h.UpdateTask).Methods("PUT")
	router.HandleFunc(commonPath, h.DeleteTask).Methods("DELETE")
}
