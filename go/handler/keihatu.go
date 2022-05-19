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
		Keihatu *response.Keihatu `json:"keihatu"`
	}
)

func NewKeihatu() IKeihatu {
	return &Keihatu{
		keihatuService: service.NewKeihatu(),
	}
}

func (h *Keihatu) Index(c echo.Context) error {

	keihatu, err := h.keihatuService.GetKeihatuByID()
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	return c.JSON(200, &JSONKeihatuIndex{
		Keihatu: response.NewKeihatu(keihatu),
	})
}
