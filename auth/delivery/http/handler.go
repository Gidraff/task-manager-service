package http

import (
	"encoding/json"
	"errors"
	"github.com/Gidraff/task-manager-service/auth"
	"github.com/Gidraff/task-manager-service/pkg/utils/helpers"
	"log"
	"net/http"
)

// AuthHandler represent http handler for user
type JsonInput struct {
	Username string
	Email    string
	Password string
}

type AuthHandler struct {
	authUC auth.UseCase
}

// NewUserHandler will initialize user resources endpoint
func NewAuthHandler(uc auth.UseCase) *AuthHandler {
	return &AuthHandler{uc}
}

// SignUp handler registers a new user
func (ah AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var user JsonInput
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	defer r.Body.Close()
	if err := decoder.Decode(&user); err != nil {
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
	hashPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		log.Printf("Failed to hash password %s", err)
	}
	err = ah.authUC.RegisterAccount(user.Username, user.Email, hashPassword)
	if err != nil {
		helpers.ErrorResponse(409, "Sign up failed. Email should be unique.", w)
		return
	}
	helpers.SuccessResponse(helpers.Message(true, "Successfully registered"), w)
	return
}
