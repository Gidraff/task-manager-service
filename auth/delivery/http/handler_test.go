package http

import (
	"errors"
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Gidraff/task-manager-service/auth/usecase"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthHandler_SignUp(t *testing.T) {
	// Given this
	type userData struct {
		username string
		email    string
		password string
	}

	type test struct {
		name     string
		input    string
		expected string
		code     int
	}

	tt := []test{

		{
			name:     "test valid json input",
			input:    `{"username":"janedoe", "email": "jane2@gmail.com", "password":"1234qwert"}`,
			expected: `{"message":"User successfully registered."}`,
			code:     201,
		},
		{
			name:     "test invalid user email",
			input:    `{"username":"janedoe", "email": "janemail.com", "password":"1234qwert"}`,
			expected: `{"message":"Invalid email format."}`,
			code:     400,
		},
	}

	r := mux.NewRouter()
	uc := new(usecase.AuthUseCaseMock)
	RegisterHandler(r, uc)

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			uc.On("GetUserByEmail", mock.Anything).Return(nil)
			uc.On("Register", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
			req := httptest.NewRequest("POST", "/api/v1/auth/sign-up", strings.NewReader(tc.input))
			res := httptest.NewRecorder()

			// When
			r.ServeHTTP(res, req)
			// Then
			assert.Equal(t, tc.code, res.Code)
			assert.Equal(t, tc.expected, res.Body.String())
			uc.AssertExpectations(t)
		})
	}
}

func TestSignIn(t *testing.T) {
	var testEmail = "testuser12@smail.com"
	// Given
	authData := `{"email": "testuser12@smail.com", "password":"45678jkj"}`

	//expected := `{"message":"Invalid email format.","token":"adkfhkdhrbfbskdfbsgadcbd"}`
	r := mux.NewRouter()
	uc := new(usecase.AuthUseCaseMock)
	RegisterHandler(r, uc)

	uc.On("GetUserByEmail", testEmail).Return(nil, errors.New("record not found"))
	req := httptest.NewRequest("POST", "/api/v1/auth/sign-in", strings.NewReader(authData))
	res := httptest.NewRecorder()

	// When
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code)
	fmt.Println(res.Body.String())
	//assert.Equal(t, expected, res.Body.String())
	uc.AssertExpectations(t)
}
