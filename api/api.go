package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ApiRouting(e *echo.Echo) {

	api := e.Group("/api")
	api.Use(middleware.Logger())
	api.Use(middleware.Recover())
	api.Use(middleware.CORS())

	api.GET("/ping", ping)

	api.GET("/time", getTime)

	api.POST("/echo", echoMsg)
}
