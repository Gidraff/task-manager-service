package http

import (
	"github.com/Gidraff/task-manager-service/auth"
	"github.com/gorilla/mux"
)

// RegisterHandler register handlers
func RegisterHandler(router *mux.Router, uc auth.UseCase) {
	h := NewAuthHandler(uc)
	router.HandleFunc("/api/v1/auth/sign-up", h.SignUp)
	router.HandleFunc("/api/v1/auth/sign-in", h.SignIn)
}
