package http

import (
	"github.com/Gidraff/task-manager-service/auth"
	"github.com/gorilla/mux"
)

func RegisterHttpEndpoints(router *mux.Router, uc auth.BasicAuthUseCase) {
	h := NewAuthHandler(uc)
	router.HandleFunc("/api/v1/auth/signup", h.SignUp)
}
