package handler

import (
	"app/service"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type (
	IEvolution interface {
		Index(c echo.Context) error
	}

	Evolution struct {
		evolutionService service.IEvolution
	}
)

func NewEvolution() IEvolution {
	return &Evolution{
		evolutionService: service.NewEvolution(),
	}
}

func (h *Evolution) Index(c echo.Context) error {
	queryId := c.QueryParam("id")
	if len(queryId) == 0 {
		return c.JSON(http.StatusBadRequest, "id is required")
	}
	userId, err := strconv.ParseInt(queryId, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "id is invalid")
	}
	status, err := h.evolutionService.GetStatus(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	return c.JSON(200, status)
}
