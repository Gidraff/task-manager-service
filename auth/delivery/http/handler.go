package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Gidraff/task-manager-service/auth"
	"github.com/Gidraff/task-manager-service/pkg/utils/helpers"
)

const (
	// ServerErrorMessage http error
	ServerErrorMessage = "Something went wrong while processing the request"
)

// JSONInput represent signup request for user
type JSONInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SignInRequest represent sign in request
type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AuthHandler encapsulate usecase
type AuthHandler struct {
	authUC auth.UseCase
}

// NewAuthHandler will initialize user resources endpoint
func NewAuthHandler(uc auth.UseCase) *AuthHandler {
	return &AuthHandler{uc}
}

// SignUp handler registers a new user
func (ah AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var account JSONInput
	var msg string

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	defer r.Body.Close()
	if err := decoder.Decode(&account); err != nil {
		msg := "Request body contains badly-formatted JSON (at position %d)"
		helpers.Response(http.StatusBadRequest, msg, w)
		return
	}
	isValid, err := helpers.SignUpValidateInput(account.Username, account.Email, account.Password)
	if !isValid {
		helpers.Response(http.StatusBadRequest, err.Error(), w)
		return
	}
	user, _ := ah.authUC.GetUserByEmail(account.Email)
	if user != nil {
		msg = "Email is taken. Try again"
		helpers.Response(http.StatusConflict, msg, w)
		return
	}
	hashPassword, err := helpers.HashPassword(account.Password)
	if err != nil {
		log.Printf("Failed to hash password %s", err)
	}
	err = ah.authUC.Register(account.Username, account.Email, hashPassword)
	if err != nil {
		msg = "Something went wrong while processing your request."
		helpers.Response(http.StatusInternalServerError, msg, w)
		return
	}
	msg = "User successfully registered."
	helpers.Response(http.StatusCreated, msg, w)
	return
}

// SignIn handler handles sign ip request
func (ah AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var signInInfo SignInRequest
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	if err := decoder.Decode(&signInInfo); err != nil {
		msg := "Request body contains badly-formatted JSON (at position %d)"
		helpers.Response(http.StatusBadRequest, msg, w)
		return
	}
	isValid, err := helpers.AuthValidateInput(signInInfo.Email, signInInfo.Password)
	if !isValid {
		helpers.Response(http.StatusBadRequest, err.Error(), w)
		return
	}

	user, err := ah.authUC.GetUserByEmail(signInInfo.Email)
	if err != nil {
		helpers.Response(http.StatusConflict, "Email not registered", w)
		return
	}

	if !helpers.ComparePassword(signInInfo.Password, user.Password) {
		helpers.Response(http.StatusUnauthorized, "Unauthorised", w)
		return
	}
	token, err := helpers.GenerateToken(user.ID)
	if err != nil {
		helpers.Response(http.StatusInternalServerError, ServerErrorMessage, w)
		return
	}
	helpers.AuthResponse(http.StatusOK, token, w)
	return
}
