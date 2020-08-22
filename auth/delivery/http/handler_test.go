package http

import (
	"github.com/Gidraff/task-manager-service/auth/usecase"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http/httptest"
	"strings"
	"testing"
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
			input:    `{"username":"janedoe", "email": "jane@mail.com", "password":"1234qwert"}`,
			expected: `{"message":"Successfully registered","status":"success"}`,
			code:     201,
		},
		{
			name:     "test invalid user email",
			input:    `{"username":"janedoe", "email": "janemail.com", "password":"1234qwert"}`,
			expected: `{"message":"Invalid email format.","status":false}`,
			code:     400,
		},
	}

	r := mux.NewRouter()
	uc := new(usecase.AuthUseCaseMock)
	RegisterHttpEndpoints(r, uc)

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			uc.On("RegisterAccount", mock.Anything, mock.Anything, mock.Anything).Return(nil)
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
