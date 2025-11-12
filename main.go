package main

import (
	"github.com/blancnoir256/Portfolio-V2-application/api"
	"github.com/blancnoir256/Portfolio-V2-application/router"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api.ApiRouting(e)
	router.Routing(e)

	e.Logger.Fatal(e.Start(":8080"))
}
