package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong!")
}

func getTime(c echo.Context) error {
	return c.String(http.StatusOK, time.Now().Format(time.RFC3339))
}

func echoMsg(c echo.Context) error {
	req := new(RequestPostEcho)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "echo endpoint requires JSON body like {\"msg\":\"...\"}")
	}
	if req.Msg == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "field \"msg\" is required")
	}
	return c.String(http.StatusOK, req.Msg)
}
