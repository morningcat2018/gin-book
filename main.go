package main

import (
	"github.com/gin-gonic/gin"
	"github.com/morningcat2018/book-service/router"
)

func main() {
	r := gin.Default()

	// 路径练习
	router.PingRouter(r)
	router.BookRouter(r)

	err := r.Run() // 监听并在 0.0.0.0:8080 上启动服务
	if err != nil {
		println("start error", err)
	}
}
