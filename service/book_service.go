package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Add(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"code": "0",
	})
}
