package form

import (
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
)

type (
	Calorie struct {
		Calorie int64 `json:"calorie" valid:"required~user calorie を入力してください"`
	}
)

func NewCalorie(c echo.Context) (*Calorie, error) {
	f := &Calorie{}

	if err := c.Bind(f); err != nil {
		return nil, err
	}

	if _, err := govalidator.ValidateStruct(f); err != nil {
		return nil, err
	}

	return f, nil
}
