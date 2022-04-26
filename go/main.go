package main

import (
	"app/db"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type Users struct {
	ID   int64
	Name string
}

func main() {
	users := []Users{}
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		db, err := db.GetSession()
		if err != nil {
			return c.String(http.StatusBadRequest, fmt.Sprintf("%v", err))
		}
		db.Select("*").From("Users").Load(&users)
		if len(users) == 0 {
			return c.String(http.StatusOK, "data is not found")
		}
		return c.String(http.StatusOK, fmt.Sprintf("%v", users))
	})
	e.Logger.Fatal(e.Start(":8080"))
}
