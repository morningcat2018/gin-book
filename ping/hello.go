package ping

import (
	"github.com/gin-gonic/gin"
	"github.com/morningcat2018/book-service/entity"
	"net/http"
)

func Hello(c *gin.Context) {
	data := entity.Book{BookCode: "G10001", BookName: "Effective Golang", Price: 8750}
	c.JSON(http.StatusOK, data)
}

/**
query string
*/
func GetById(c *gin.Context) {
	bookNo := c.Query("bookNo")
	bookNo2, ok := c.GetQuery("bookNo2")
	if ok {
		c.JSON(http.StatusOK, gin.H{"bookNo2": bookNo2})
	} else {
		c.JSON(http.StatusOK, gin.H{"bookNo": bookNo})
	}
}

/**
form 表单
*/
func FormById(c *gin.Context) {
	userNo := c.PostForm("userNo")
	password := c.PostForm("password")
	c.JSON(http.StatusOK, gin.H{"userNo": userNo, "password": password})
}

/**
path 路径
*/
func Path(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{"name": name})
}

/**
json 参数
*/
func Json(c *gin.Context) {
	var b entity.Book
	err := c.ShouldBind(&b)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusOK, b)
	}

}
