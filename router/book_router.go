package router

import (
	"github.com/gin-gonic/gin"
	"github.com/morningcat2018/book-service/service"
)

func BookRouter(r *gin.Engine) {

	bookGroup := r.Group("/book")
	{
		bookGroup.POST("/add", service.BookAdd)
		bookGroup.GET("/getDetail", service.BookAdd)
		bookGroup.POST("/page", service.BookAdd)
		bookGroup.GET("/delete", service.BookAdd)
	}

}
