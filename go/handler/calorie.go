package handler

import (
	"app/form"
	model "app/model"
	"app/response"
	"app/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type (
	ICalorie interface {
		Index(c echo.Context) error
		Create(c echo.Context) error
	}

	Calorie struct {
		calorieService service.ICalorie
	}

	JSONCalorieIndex struct {
		Calories *response.Calories `json:"calories"`
	}
)

func NewCalorie() ICalorie {
	return &Calorie{
		calorieService: service.NewCalorie(),
	}
}

func (h *Calorie) Index(c echo.Context) error {
	queryId := c.QueryParam("id")
	if len(queryId) == 0 {
		return c.JSON(http.StatusBadRequest, "id is required")
	}
	queryCalorieType := c.QueryParam("calorie_type")
	if len(queryCalorieType) == 0 {
		return c.JSON(http.StatusBadRequest, "calorie_type is required")
	}
	userId, err := strconv.ParseInt(queryId, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "id is invalid")
	}
	calorieTypeInt64, err := strconv.ParseInt(queryCalorieType, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "calorie_type is invalid")
	}
	calorieType := model.CalorieType(calorieTypeInt64)
	if calorieType.IsInvalid() {
		return c.JSON(http.StatusBadRequest, "calorie_type is invalid, please select 0 or 1")
	}
	calories, err := h.calorieService.GetCaloriesByUserIdCalorieType(userId, calorieType)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("%v", err))
	}
	if len(*calories) == 0 {
		return c.JSON(http.StatusBadRequest, "calories is not found")
	}

	return c.JSON(200, &JSONCalorieIndex{
		Calories: response.NewCalories(calories),
	})
}

func (h *Calorie) Create(c echo.Context) error {
	queryId := c.QueryParam("id")
	if len(queryId) == 0 {
		return c.JSON(http.StatusBadRequest, "id is required")
	}
	queryCalorieType := c.QueryParam("calorie_type")
	if len(queryCalorieType) == 0 {
		return c.JSON(http.StatusBadRequest, "calorie_type is required")
	}
	userId, err := strconv.ParseInt(queryId, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "id is invalid")
	}
	calorieTypeInt64, err := strconv.ParseInt(queryCalorieType, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "calorie_type is invalid")
	}
	calorieType := model.CalorieType(calorieTypeInt64)
	if calorieType.IsInvalid() {
		return c.JSON(http.StatusBadRequest, "calorie_type is invalid, please select 0 or 1")
	}
	f, err := form.NewCalorie(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	if err := h.calorieService.Create(f, userId, calorieType); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	return c.JSON(201, nil)
}
