package postgres

import (
	"database/sql"
	"github.com/Gidraff/task-manager-service/auth"
	"log"
	"time"

	"github.com/Gidraff/task-manager-service/model"
)

type userRepo struct {
	Db *sql.DB
}

// NewAuthRepo returns a new UserRepository interface
func NewUserRepo(db *sql.DB) auth.UserRepository {
	return &userRepo{Db: db}
}

// CreateUser persists a new user to db
func (ur *userRepo) CreateUser(u *model.User) error {
	query := "INSERT INTO users (username,email,password,created_at) VALUES ($1,$2,$3,$4)"
	stmt, err := ur.Db.Prepare(query) // here context is used for the preparation of the statement
	if err != nil {
		return err
	}

	res, err := stmt.Exec(
		u.Username, u.Email, u.Password, time.Now())

	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		log.Fatalf("Expected to affect 1 row, affected %d", rows)
	}

	return nil
}

// fetch returns a list of users or error
func (ur *userRepo) fetch(query string, args ...interface{}) ([]*model.User, error) {
	rows, err := ur.Db.Query(query, args)
	if err != nil {
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Printf("Error %s", err)
		}
	}()

	result := make([]*model.User, 0)
	for rows.Next() {
		u := new(model.User)
		err = rows.Scan(
			&u.ID,
			&u.Username,
			&u.Email,
			&u.CreatedAt,
		)
		if err != nil {
			log.Printf("Error on fetch %s", err)
			return nil, err
		}
		result = append(result, u)
	}
	return result, nil

}

// GetUserByEmail returns a user with the provided email
func (ur *userRepo) GetUserByEmail(email string) (res *model.User, err error) {
	query := `SELECT id, username, email FROM users WHERE id=$1`

	list, err := ur.fetch(query, email)
	if err != nil {
		return
	}
	if len(list) > 0 {
		res = list[0]
	} else {
		return res, err
	}
	return
}
