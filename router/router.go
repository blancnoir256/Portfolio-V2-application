package router

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func Routing(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "Pong!")
	})

	e.GET("/time", func(c echo.Context) error {
		return c.String(http.StatusOK, time.Now().Format(time.RFC3339))
	})

	e.POST("/echo", func(c echo.Context) error {
		req := new(RequestPostEcho)
		if err := c.Bind(req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "echo endpoint requires JSON body like {\"msg\":\"...\"}")
		}
		if req.Msg == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "field \"msg\" is required")
		}
		return c.String(http.StatusOK, req.Msg)
	})
}
