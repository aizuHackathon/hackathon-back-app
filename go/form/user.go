package form

import (
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
)

type (
	User struct {
		// ID     int64   `json:"id"  valid:"required~user ID を入力してください"`
		Name   string  `json:"name" valid:"required~user name を入力してください"`
		Height float32 `json:"height" valid:"required~user height を入力してください"`
		Weight float32 `json:"weight" valid:"required~user weight を入力してください"`
		Sex    int64   `json:"sex" valid:"required~user sex を入力してください"`
		Old    int64   `json:"old" valid:"required~user old を入力してください"`
	}
)

func NewUser(c echo.Context) (*User, error) {
	f := &User{}

	if err := c.Bind(f); err != nil {
		return nil, err
	}

	if _, err := govalidator.ValidateStruct(f); err != nil {
		return nil, err
	}

	return f, nil
}
