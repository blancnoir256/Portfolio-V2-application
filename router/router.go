package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Routing(e *echo.Echo, distDir string) {

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", distDir)

}
