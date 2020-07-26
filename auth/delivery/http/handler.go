package http

import (
	"encoding/json"
	"errors"
	"github.com/Gidraff/task-manager-service/auth"
	"github.com/Gidraff/task-manager-service/model"
	"github.com/Gidraff/task-manager-service/pkg/utils/helpers"
	"net/http"
)

// AuthHandler represent http handler for user
type AuthHandler struct {
	authUC auth.UseCase
}

// NewUserHandler will initialize user resources endpoint
func NewAuthHandler(uc auth.UseCase) *AuthHandler {
	return &AuthHandler{uc}
}

// SignUp handler registers a new user
func (ah AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var user model.User
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	defer r.Body.Close()

	if  err := decoder.Decode(&user); err != nil {
		var syntaxError *json.SyntaxError
		switch {
		case errors.As(err, &syntaxError):
			msg := "Request body contains badly-formatted JSON (at position %d)"
			helpers.ErrorResponse(400, msg, w)
			return
		default:
			return
		}
	}
	isValid, message := helpers.IsUserInfoValid(user.Username, user.Email, user.Password)
	if !isValid {
		helpers.ErrorResponse(400, message, w)
		return
	}

	duplicateUser, err := ah.authUC.GetUserByEmail(user.Email)
	if  duplicateUser != nil {
		message := "User with that email already exist."
		helpers.ErrorResponse(http.StatusBadRequest, message, w)
		return
	}

	err = ah.authUC.SignUp(&user)
	if err != nil {
		helpers.ErrorResponse(http.StatusInternalServerError, "duplicate", w)
		return
	}
	helpers.SuccessResponse(helpers.Message(true, "Successfully registered"), w)
	return
}
