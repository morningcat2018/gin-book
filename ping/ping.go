package ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
ping 函数
*/
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func NotFount(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": "请求地址有误",
	})
}
