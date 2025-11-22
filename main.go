package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/blancnoir256/Portfolio-V2-application/api"
	"github.com/blancnoir256/Portfolio-V2-application/router"
	"github.com/labstack/echo"
)

func main() {

	// 環境変数 APP_PORT を取得
	port := os.Getenv("APP_PORT")
	if port == "" {
		fmt.Println("環境変数 APP_PORT が設定されていません。デフォルトポート 8080 を使用します。")
		// デフォルトポート
		port = "8080"
	}

	// ポート番号を整数に変換
	portNum, err := strconv.Atoi(port)
	if err != nil {
		fmt.Printf("無効なポート番号が指定されています: %s\n", port)
		os.Exit(1)
	}

	distDir := os.Getenv("DIST_DIR")
	if distDir == "" {
		fmt.Println("環境変数 DIST_DIR が設定されていません。デフォルトディレクトリ /dist を使用します。")
		// デフォルトディレクトリ
		distDir = "/dist"
	}

	e := echo.New()

	api.ApiRouting(e)
	router.Routing(e, distDir)

	// サーバーを起動
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", portNum)))
}
