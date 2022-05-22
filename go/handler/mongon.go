package handler

import (
	"app/response"
	"app/service"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type (
	IMongon interface {
		Index(c echo.Context) error
	}

	Mongon struct {
		mongonService service.IMongon
	}

	JSONMongonIndex struct {
		Mongon *response.Mongon `json:"mongon"`
	}
)

func NewMongon() IMongon {
	return &Mongon{
		mongonService: service.NewMongon(),
	}
}

func (h *Mongon) Index(c echo.Context) error {
	mongon, err := h.mongonService.GetMongon()
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	return c.JSON(200, &JSONMongonIndex{
		Mongon: response.NewMongon(mongon),
	})
}
