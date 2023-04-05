package entity

import (
	"github.com/morningcat2018/book-service/model"
	"gorm.io/gorm"
	"reflect"
)

type Book struct {
	gorm.Model
	//Id          int64      `gorm:"AUTO_INCREMENT" json:"id"`
	BookCode    string           `gorm:"type:varchar(10)" json:"bookCode"`
	BookName    string           `gorm:"type:varchar(50)" json:"bookName"`
	Author      string           `gorm:"type:varchar(25)" json:"author"`
	PublishDate *model.LocalDate `json:"publishDate"`
	Price       int              `gorm:"size:3" json:"price"`
	//LogicStatus int        `gorm:"size:3" json:"logicStatus"`
}

// 自定义表名
func (Book) TableName() string {
	return "gin_book"
}

func (s Book) IsEmpty() bool {
	return reflect.DeepEqual(s, Book{})
}

//type User struct {
//	gorm.Model
//	Name          string `gorm:"type:varchar(20)"`
//	Password      string `gorm:"type:varchar(20)"`
//	LastLoginTime *model.LocalTime
//}
