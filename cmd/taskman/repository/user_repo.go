package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/lib/pq"
	"log"
	"time"

	"github.com/Gidraff/task-manager-service/cmd/taskman/model"
)

// UserRepository encapsulates the logic to access user from the data source
type UserRepository interface {
	Create(ctx context.Context, u *model.User) error
	//GetByEmail(ctx context.Context, email string) (error *model.User)
}

type userRepo struct {
	Conn *sql.DB
}

// NewUserRepo returns a new UserRepository interface
func NewUserRepo(db *sql.DB) UserRepository {
	return &userRepo{Conn: db}
}

// Create adds user to database
func (ur *userRepo) Create(ctx context.Context, u *model.User) error {
	query := "INSERT INTO users (username,email,password,created_at) VALUES ($1,$2,$3,$4)"
	stmt, err := ur.Conn.PrepareContext(ctx, query) // here context is used for the preparation of the statement
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(
		ctx, u.Username, u.Email, u.Password, time.Now())
	if pgerr, ok := err.(*pq.Error); ok {
		if pgerr.Code == "23505" {
			fmt.Println("Before error===>")
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

func (ur *userRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*model.User, error) {
	rows, err := ur.Conn.QueryContext(ctx, query, args)
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
