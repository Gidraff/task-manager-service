package http

import (
	"encoding/json"
	"github.com/Gidraff/task-manager-service/auth"
	"github.com/Gidraff/task-manager-service/pkg/utils"
	"net/http"

	"github.com/Gidraff/task-manager-service/model"
)

// AuthHandler represent http handler for user
type AuthHandler struct {
	basicAuth auth.BasicAuthUseCase
}

// NewUserHandler will initialize user resources endpoint
func NewAuthHandler(uc auth.BasicAuthUseCase) *AuthHandler {
	return &AuthHandler{uc}
}

// SignUp handler registers a new user
func (ah AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var input model.User
	err := decoder.Decode(&input)

	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	err = ah.basicAuth.SignUp(&input)
	if err != nil {
		utils.ErrorResponse(http.StatusInternalServerError, "duplicate", w)
		return
	}
	utils.SuccessResponse(utils.Message(true, "Successfully created!"), w)
}
