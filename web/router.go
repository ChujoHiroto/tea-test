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
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowWebSockets = true
	corsConfig.AllowCredentials = true
	r.Use(cors.New(corsConfig))

	// 各種ハンドラーの生成
	helloHandler := handler.NewHelloHandler("Hello World")

	// Routing
	r.GET("/", helloHandler.HandleGetHello)
	r.POST("/", helloHandler.HandlePostHello)

	return r
}
