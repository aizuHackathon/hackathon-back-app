package handler

import (
	"app/response"
	"app/service"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type (
	IKeihatu interface {
		Index(c echo.Context) error
	}

	Keihatu struct {
		keihatuService service.IKeihatu
	}

	JSONKeihatuIndex struct {
		Keihatus *response.Keihatus `json:"keihatus"`
	}
)

func NewKeihatu() IKeihatu {
	return &Keihatu{
		keihatuService: service.NewKeihatu(),
	}
}

func (h *Keihatu) Index(c echo.Context) error {
	keihatus, err := h.keihatuService.GetAllKeihatus()
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("%v", err))
	}
	if len(*keihatus) == 0 {
		return c.JSON(http.StatusBadRequest, "keihatu is not found")
	}

	return c.JSON(200, &JSONKeihatuIndex{
		Keihatus: response.NewKeihatus(keihatus),
	})
}
