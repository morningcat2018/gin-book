package router

import (
	"github.com/gin-gonic/gin"
	"github.com/morningcat2018/book-service/ping"
)

func PingRouter(r *gin.Engine) {
	// 路由
	r.GET("/ping", ping.Ping)
	r.GET("/hello", ping.Hello)
	r.GET("/find", ping.GetById)
	r.POST("/form", ping.FormById)

	/**
	  /name/Tony
	*/
	r.GET("/name/:name", ping.Path)

	/**
	  {
	  	"bookNo": "A1001",
	  	"bookName": "golang in action",
	  	"price": 1001
	  }
	*/
	r.POST("/json", ping.Json)

	r.Any("/any", ping.Ping)

	/**
	访问无效路径
	*/
	r.NoRoute(ping.NotFount)
}
