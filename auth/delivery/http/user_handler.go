package http

import (
	"encoding/json"
	"github.com/Gidraff/task-manager-service/auth"
	"github.com/Gidraff/task-manager-service/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/Gidraff/task-manager-service/model"
)

// UserHandler represent httphandler for user
type AuthHandler struct {
	useCase auth.UseCase
}

// NewUserHandler will initialize user resources endpoint
func NewUserHandler(router *mux.Router, uc auth.UseCase) {
	authHandler := &AuthHandler{uc}

	router.HandleFunc("/api/v1/auth/signup", authHandler.Signup)
}

// Signup will sign up the user by given req body
func (ah AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var input model.User
	err := decoder.Decode(&input)
	if err != nil {

		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}
	err = ah.useCase.Register(&input)
	if err != nil {
		utils.ErrorResponse(http.StatusInternalServerError, "duplicate", w)
		return
	}
	utils.SuccessResponse(utils.Message(true, "Successfully created!"), w)
}
