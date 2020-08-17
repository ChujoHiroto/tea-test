package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// / のエンドポイントを管理する構造体です。
type HelloHandler struct {
	message string
}

// HealthHandlerのポインタを生成する関数です。
func NewHelloHandler(message string) *HelloHandler {
	return &HelloHandler{message: message}
}

// GETリクエストに対応した関数
func (h *HelloHandler) HandleGetHello(c *gin.Context) {
	res := &helloRes{
		Message: h.message,
	}

	c.JSON(http.StatusOK, res)
}

// POSTリクエストに対応した関数
func (h *HelloHandler) HandlePostHello(c *gin.Context) {
	// リクエスト内容を定義して、JSONとGOの構造体を対応させる。
	// {
	//	   "message": "Hello!!!"
	// }
	type helloRequest struct {
		Message string `json:"message"`
	}

	requestBody := helloRequest{}
	c.BindJSON(&requestBody)

	// サーバー内のメッセージ内容を書き換える！
	h.message = requestBody.Message

	c.Status(http.StatusOK)
}

type helloRes struct {
	Message string `json:"message"`
}
