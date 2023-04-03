package router

import (
	"github.com/gin-gonic/gin"
	"github.com/morningcat2018/book-service/service"
)

func BookRouter(r *gin.Engine) {

	bookGroup := r.Group("/book")
	{
		bookGroup.POST("/add", service.Add)
		bookGroup.GET("/getDetail", service.Add)
		bookGroup.POST("/page", service.Add)
		bookGroup.GET("/delete", service.Add)
	}

}
