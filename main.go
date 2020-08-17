package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/chujohiroto/tea-test/web"
)

func main() {
	// WebServerの生成
	s := web.NewServer()

	// シグナルを受け取れるようにgoroutine内でサーバを起動する
	go func() {
		port := os.Getenv("PORT")
		if port == "" {
			port = "8888"
		}

		if err := s.Run(":" + port); err != nil {
			fmt.Printf("shutting down the server with error: %v", err)
			os.Exit(1)
		}
	}()

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	fmt.Printf("SIGNAL %d received, then shutting down...", <-quit)
}
