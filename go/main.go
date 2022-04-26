package main

import (
	"app/router"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	router.Init(e)
}
