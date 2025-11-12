package api

import (
	"github.com/labstack/echo"
)

func ApiRouting(e *echo.Echo) {

	e.GET("/api/ping", ping)

}
