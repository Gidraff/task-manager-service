package http

import (
	"github.com/Gidraff/task-manager-service/project"
	//"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

// RegisterHandler registers hand;ers
func RegisterHandler(router *mux.Router, uc project.IProjectUsecase) {
	h := NewProjectHandler(uc)
	router.HandleFunc("/api/v1/project/", h.CreateProject).Methods("POST")
	router.HandleFunc("/api/v1/project/", h.GetAllProjects).Methods("GET")
	router.HandleFunc("/api/v1/project/{id:[0-9]+}", h.GetProjectByID).Methods("GET")
	router.HandleFunc("/api/v1/project", h.GetProjectByName).Methods("GET").Queries("name", "{name}")
	router.HandleFunc("/api/v1/project/{id}", h.UpdateProject).Methods("PUT")
	router.HandleFunc("/api/v1/project/{id}", h.DeleteProject).Methods("DELETE")
}
