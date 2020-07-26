package postgres

import (
	"database/sql/driver"
	"regexp"

	//"errors"
	"github.com/jinzhu/gorm"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Gidraff/task-manager-service/model"
	"github.com/stretchr/testify/assert"
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
	gdb, err := gorm.Open("postgres", db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()
	const sqlInsert = `INSERT INTO "users" ("created_at","updated_at","deleted_at","username","email","password","status") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "users"."id"`
	const newId = 1
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(sqlInsert)).
		WithArgs(AnyTime{}, AnyTime{}, AnyTime{}, u.Username, u.Email, u.Password, false).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).AddRow(newId))
	mock.ExpectCommit() // commit transaction

	// Behaviour to be tested
	userRepo := NewUserRepo(gdb)
	err = userRepo.Store(u)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		log.Fatalf("unfulfilled expectations: %s", err)
	}
}
