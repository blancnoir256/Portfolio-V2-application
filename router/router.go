package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Routing(e *echo.Echo, distDir string) {

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", distDir)

}
