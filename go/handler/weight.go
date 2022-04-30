package handler

import (
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
	user_id, err := strconv.ParseInt(c.QueryParam("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "id is not int")
	}
	weights, err := h.weightService.GetAllWeights(int64(user_id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("%v", err))
	}
	if len(*weights) == 0 {
		return c.JSON(http.StatusBadRequest, "weights is not found")
	}

	return c.JSON(200, &JSONWeightIndex{
		Weights: response.NewWeights(weights),
	})
}
