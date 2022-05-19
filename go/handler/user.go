package handler

import (
	"app/form"
	"app/response"
	"app/service"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type (
	IUser interface {
		Index(c echo.Context) error
		CheckName(c echo.Context) error
		Login(c echo.Context) error
		Create(c echo.Context) error
	}

	User struct {
		userService service.IUser
	}

	JSONUserIndex struct {
		Users *response.Users `json:"users"`
	}

	JSONUserID struct {
		ID *int64 `json:"id"`
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

func (h *User) CheckName(c echo.Context) error {
	f, err := form.NewUserName(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	count, err := h.userService.GetNameCount(f)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("%v", err))
	}
	if count != int64(0) {
		return c.JSON(http.StatusBadRequest, "name is used")
	}

	return c.JSON(201, nil)
}

func (h *User) Login(c echo.Context) error {
	name := c.QueryParam("name")
	if len(name) == 0 {
		return c.JSON(http.StatusBadRequest, "name is required")
	}
	pass := c.QueryParam("pass")
	if len(pass) == 0 {
		return c.JSON(http.StatusBadRequest, "pass is required")
	}

	id, err := h.userService.GetUserIdByNamePass(name, pass)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	return c.JSON(201, &JSONUserID{
		ID: id,
	})
}

func (h *User) Create(c echo.Context) error {
	f, err := form.NewUser(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	if err := h.userService.Create(f); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	return c.JSON(201, nil)
}
