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
	// goroutineについては → https://go-tour-jp.appspot.com/concurrency/1
	go func() {
		// 環境変数 or シェル変数を使ってPortを指定します！
		// 環境変数 or シェル変数については → https://qiita.com/chihiro/items/bb687903ee284766e879
		// Portについてはこちら → https://ja.wikipedia.org/wiki/%E3%83%9D%E3%83%BC%E3%83%88_(%E3%82%B3%E3%83%B3%E3%83%94%E3%83%A5%E3%83%BC%E3%82%BF%E3%83%8D%E3%83%83%E3%83%88%E3%83%AF%E3%83%BC%E3%82%AF)#:~:text=%E3%82%B3%E3%83%B3%E3%83%94%E3%83%A5%E3%83%BC%E3%82%BF%E3%83%8D%E3%83%83%E3%83%88%E3%83%AF%E3%83%BC%E3%82%AF%E3%81%AB%E3%81%8A%E3%81%84%E3%81%A6%E3%80%81%E3%83%9D%E3%83%BC%E3%83%88%EF%BC%88port,%E3%81%AE%E3%82%A8%E3%83%B3%E3%83%89%E3%83%9D%E3%82%A4%E3%83%B3%E3%83%88%E3%81%A7%E3%81%82%E3%82%8B%E3%80%82&text=%E3%83%9D%E3%83%BC%E3%83%88%E3%81%AF%E3%83%9B%E3%82%B9%E3%83%88%E3%81%AEIP,%E3%81%A8%E5%AE%9B%E5%85%88%E3%81%8C%E6%B1%BA%E5%AE%9A%E3%81%99%E3%82%8B%E3%80%82
		port := os.Getenv("PORT")
		if port == "" {
			port = "8889"
		}

		if err := s.Run(":" + port); err != nil {
			fmt.Printf("shutting down the server with error: %v", err)
			os.Exit(1)
		}
	}()

	// シグナルを受け取ってシャットダウンする
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	fmt.Printf("SIGNAL %d received, then shutting down...\n", <-quit)
}
