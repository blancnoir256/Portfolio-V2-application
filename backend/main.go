package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	appDir    = ""
	secretDir = ""
	isDev     = false
)

func main() {
	commandLineArguments()
	envLoadAndSet()

	e := echo.New()
	setupMiddleware(e)
	setupRoutes(e)
	startServer(e)
}

// ミドルウェアを設定する関数
func setupMiddleware(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}

// ルーティングを設定する関数
// pingと静的ファイルのルーティングとSPAのフォールバックを設定
func setupRoutes(e *echo.Echo) {
	setupAPIRoutes(e)
	setupStaticRoutes(e)

	e.GET("/*", spaFallbackHandler)
}
