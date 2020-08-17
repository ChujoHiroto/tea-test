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

func (h *HelloHandler) HandleGetHello(c *gin.Context) {
	res := &helloRes{
		Message: h.message,
	}

	c.JSON(http.StatusOK, res)
}

func (h *HelloHandler) HandlePostHello(c *gin.Context) {
	type helloRequest struct {
		Message string `json:"message"`
	}
	requestBody := helloRequest{}
	c.BindJSON(&requestBody)

	h.message = requestBody.Message

	c.Status(http.StatusOK)
}

type helloRes struct {
	Message string `json:"message"`
}
