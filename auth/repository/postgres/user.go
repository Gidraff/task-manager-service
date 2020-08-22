package postgres

import (
	"database/sql"
	"github.com/Gidraff/task-manager-service/auth"
	"github.com/Gidraff/task-manager-service/model"
	"github.com/getsentry/sentry-go"
	"github.com/lib/pq"
	"log"
	"time"
)

type userRepo struct {
	Conn *sql.DB
	//Logger *log.Logger
}

// NewAuthRepo returns a new UserRepository interface
func NewUserRepo(db *sql.DB) auth.UserRepository {
	return &userRepo{Conn: db}
}

// CreateUser persists a new user to db
func (ur *userRepo) Store(username, email, password string) error {
	query := "INSERT INTO accounts (username,email,password,active,created_on,last_login) VALUES ($1,$2,$3,$4,$5,$6)"
	stmt, err := ur.Conn.Prepare(query) // here context is used for the preparation of the statement
	if err != nil {
		return err
	}

	res, err := stmt.Exec(
		username, email, password, false, time.Now(), nil)
	if pgerr, ok := err.(*pq.Error); ok {
		if pgerr.Code == "23505" {
			sentry.CaptureException(err)
			log.Printf("Failed with %s", err)
			return err.(*pq.Error)
		}
	}
	if err != nil {
		sentry.CaptureException(err)
		log.Println(err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
	}
	if rows != 1 {
		log.Fatalf("Expected to affect 1 row, affected %d", rows)
	}
	return nil
}

func (ur *userRepo) fetch(query string, args ...interface{}) ([]*model.Account, error) {
	rows, err := ur.Conn.Query(query, args)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Printf("Error %s", err)
		}
	}()

	result := make([]*model.Account, 0)
	for rows.Next() {
		u := new(model.Account)
		err := rows.Scan(&u.ID, &u.Username, &u.Email, &u, &u.Password, u.Active)
		if err != nil {
			log.Printf("Error on fetch %s", err)
			return nil, err
		}
		result = append(result, u)
	}
	return result, nil
}

// GetUserByEmail returns a user with the provided email
func (ur *userRepo) FetchByEmail(email string) (res *model.Account, err error) {
	query := `SELECT id, username, password, email FROM accounts WHERE email=$1`
	rows, err := ur.fetch(query, email)
	if err != nil {
		return
	}
	if len(rows) > 0 {
		res = rows[0]
	} else {
		return res, err
	}
	return
}
