package http

import (
	"context"
	"log"
	"net/http"

	"github.com/Gidraff/task-manager-service/cmd/taskman/model"
	"github.com/Gidraff/task-manager-service/cmd/taskman/service"
	"github.com/labstack/echo"
)

// UserHandler represent httphandler for user
type UserHandler struct {
	UService service.Service
}

// NewUserHandler will initialize user resournces endpoint
func NewUserHandler(e *echo.Echo, us service.Service) {
	handler := &UserHandler{
		UService: us,
	}

	e.POST("/api/v1/auth/signup", handler.Signup)

}

// Signup will sign up the user by given req body
func (u UserHandler) Signup(c echo.Context) error {
	var user model.User
	err := c.Bind(&user)
	if err != nil {
		log.Printf("Http: %s", err)
		return c.JSON(http.StatusUnprocessableEntity, err)
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	err = u.UService.Register(ctx, &user)
	if err != nil {
		return c.JSON(http.StatusMethodNotAllowed, err.Error())
	}

	return c.JSON(http.StatusCreated, user)
}
