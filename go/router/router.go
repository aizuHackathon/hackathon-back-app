package router

import (
	"app/handler"
	"net/http"

	"github.com/labstack/echo"
)

type (
	IV1 interface {
		withNone(e *echo.Echo)
	}

	V1 struct {
		userHandler      handler.IUser
		weightHandler    handler.IWeight
		calorieHandler   handler.ICalorie
		evolutionHandler handler.IEvolution
		MongonHandler handler.IMongon
	}
)

func Init(e *echo.Echo) {
	var r IV1 = &V1{
		userHandler:      handler.NewUser(),
		weightHandler:    handler.NewWeight(),
		calorieHandler:   handler.NewCalorie(),
		evolutionHandler: handler.NewEvolution(),
		MongonHandler: handler.NewMongon(),
	}

	r.withNone(e)

	// 簡易版
	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello")
	})

	e.Logger.Fatal(e.Start(":8080"))
}

func (r V1) withNone(e *echo.Echo) {
	e.GET("/users", r.userHandler.Index)
	e.POST("/users", r.userHandler.Create)

	e.GET("/weight", r.weightHandler.Index)
	e.POST("/weight", r.weightHandler.Create)

	e.GET("/calorie", r.calorieHandler.Index)
	e.POST("/calorie", r.calorieHandler.Create)

	e.GET("/evolution", r.evolutionHandler.Index)

	e.GET("/name", r.userHandler.CheckName)

	e.GET("/login", r.userHandler.Login)

	e.GET("/mongon", r.MongonHandler.Index)
}
