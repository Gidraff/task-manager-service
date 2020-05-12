package repository

import (
	"context"
	"database/sql/driver"
	"log"
	"regexp"
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

func TestStore(t *testing.T) {
	u := &model.User{
		Username: "johndoe",
		Email:    "johndoe@gmail.com",
		Password: "1234qwerty",
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO users")).
		ExpectExec().
		WithArgs(u.Username, u.Email, u.Password, AnyTime{}).
		WillReturnResult(sqlmock.NewResult(1, 1))

	ctx := context.Background()
	// Behaviour to be tested
	userRepo := NewUserRepo(db)
	err = userRepo.Create(ctx, u)

	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		log.Fatalf("unfulfilled expectations: %s", err)
	}
}
