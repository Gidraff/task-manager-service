package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Gidraff/task-manager-service/models"
	"github.com/Gidraff/task-manager-service/utils"
)

// CreateAccount for a new user
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) // decode json request data into struct
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	resp := account.Create() // Create account
	utils.Respond(w, resp)
}
