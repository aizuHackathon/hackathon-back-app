package form

import (
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
)

type (
	Weight struct {
		Weight float32 `json:"weight" valid:"required~user weight を入力してください"`
	}
)

func NewWeight(c echo.Context) (*Weight, error) {
	f := &Weight{}

	if err := c.Bind(f); err != nil {
		return nil, err
	}

	if _, err := govalidator.ValidateStruct(f); err != nil {
		return nil, err
	}

	return f, nil
}
