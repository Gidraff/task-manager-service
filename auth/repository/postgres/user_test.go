package postgres

import (
	//"database/sql"
	"database/sql/driver"
	//"regexp"

	//"errors"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Gidraff/task-manager-service/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool { // implements Argument interface
	_, ok := v.(time.Time)
	return ok
}

func TestUserRepo_Store(t *testing.T) {
	u := &model.User{
		Username: "john",
		Email:    "johndoe@gmail.com",
		Password: "1234qwerty",
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	gdb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "users" (.+) RETURNING`).
		WithArgs(u.Username, u.Email, u.Password, u.CreatedAt).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	// Behaviour to be tested
	userRepo := NewUserRepo(gdb)
	err = userRepo.Store(u.Username, u.Email, u.Password)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		log.Fatalf("unfulfilled expectations: %s", err)
	}
}
