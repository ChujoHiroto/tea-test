package web

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/chujohiroto/tea-test/web/handler"
)

// ミドルウェアやハンドラーが登録されたGINのengineを返す
func NewServer() *gin.Engine {
	r := gin.Default()

	// CORSの設定
	// CORSについては→ https://qiita.com/att55/items/2154a8aad8bf1409db2b
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowWebSockets = true
	corsConfig.AllowCredentials = true
	r.Use(cors.New(corsConfig))

	// 各種ハンドラーの生成
	helloHandler := handler.NewHelloHandler("メッセージを書き換えてみてね")

	// localhost:PORT/の GETのアクセスに対応する
	r.GET("/", helloHandler.HandleGetHello)
	// localhost:PORT/の POSTのアクセスが来たら、helloHandler.HandlePostHello
	r.POST("/", helloHandler.HandlePostHello)

	return r
}
