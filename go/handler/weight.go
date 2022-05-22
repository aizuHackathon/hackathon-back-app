package handler

import (
	"app/form"
	"app/response"
	"app/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type (
	IWeight interface {
		Index(c echo.Context) error
		Create(c echo.Context) error
	}

	Weight struct {
		weightService service.IWeight
	}

	JSONWeightIndex struct {
		Weights *response.Weights `json:"weights"`
	}
)

func NewWeight() IWeight {
	return &Weight{
		weightService: service.NewWeight(),
	}
}

func (h *Weight) Index(c echo.Context) error {
	ID := c.QueryParam("id")
	if len(ID) == 0 {
		return c.JSON(http.StatusBadRequest, "id is required")
	}
	userId, err := strconv.ParseInt(ID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "id is invalid")
	}
	weights, err := h.weightService.GetWeightsByUserID(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("%v", err))
	}
	// TODO: discuss with frontend
	if len(*weights) == 0 {
		return c.JSON(http.StatusBadRequest, "weights is not found")
	}

	return c.JSON(200, &JSONWeightIndex{
		Weights: response.NewWeights(weights),
	})
}

func (h *Weight) Create(c echo.Context) error {
	ID := c.QueryParam("id")
	if len(ID) == 0 {
		return c.JSON(http.StatusBadRequest, "id is required")
	}
	userId, err := strconv.ParseInt(ID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "id is invalid")
	}
	f, err := form.NewWeight(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	if err := h.weightService.Create(f, userId); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	return c.JSON(201, nil)
}
