package service

import (
	"github.com/gin-gonic/gin"
	"github.com/morningcat2018/book-service/dao"
	"github.com/morningcat2018/book-service/entity"
	"net/http"
)

func BookAdd(c *gin.Context) {

	var book entity.Book
	c.ShouldBind(&book)
	dao.BookInsert(&book)

	c.JSON(http.StatusOK, gin.H{
		"code": "0",
		"data": book.ID,
	})
}

func BookGetById(c *gin.Context) {
	id := c.Query("id")
	book := dao.BookGetById(id)
	c.JSON(http.StatusOK, gin.H{
		"code": "0",
		"data": book,
	})
}

func BookLogicDelete(c *gin.Context) {
	id := c.Query("id")
	book := dao.BookUpdateById(id)
	c.JSON(http.StatusOK, gin.H{
		"code": "0",
		"data": book.ID,
	})
}
