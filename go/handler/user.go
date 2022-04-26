package handler

import (
	"app/response"
	"app/service"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type (
	IUser interface {
		Index(c echo.Context) error
	}

	User struct {
		userService service.IUser
	}

	JSONUserIndex struct {
		Users *response.Users `json:"users"`
	}
)

func NewUser() IUser {
	return &User{
		userService: service.NewUser(),
	}
}

func (h *User) Index(c echo.Context) error {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("%v", err))
	}
	if len(*users) == 0 {
		return c.JSON(http.StatusBadRequest, "user is not found")
	}

	return c.JSON(200, &JSONUserIndex{
		Users: response.NewUsers(users),
	})
}
