package handler

import (
	"app/response"
	"app/service"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type (
	IMongoninterface {
		Index(c echo.Context) error
	}

	Mongon struct {
		mongonService service.IMongon
	}

	JSONMongonIndex struct {
		Mongons *response.Mongons `json:"mongons"`
	}
)

func NewMongon() IMongon {
	return &Mongon{
		mongonService: service.NewMongon(),
	}
}

func (h *mongon) Index(c echo.Context) error {
	mongon, err := h.mongonService.GetAllMongon()
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("%v", err))
	}
	if len(*mongon) == 0 {
		return c.JSON(http.StatusBadRequest, "mongon is not found")
	}

	return c.JSON(200, &JSONMongonIndex{
		Mongon: response.NewMongon(mongon),
	})
}
