package auth

import (
	"github.com/Gidraff/task-manager-service/model"
)

// UserRepository encapsulates the logic to access user from the data source
type UserRepository interface {
	Store(username, email, password string) error
	FetchByEmail(email string) (model.User, error)
}
