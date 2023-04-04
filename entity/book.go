package entity

import "time"

type Book struct {
	Id          int64      `gorm:"AUTO_INCREMENT" json:"id"`
	BookCode    string     `gorm:"type:varchar(10)" json:"bookCode"`
	BookName    string     `gorm:"type:varchar(50)" json:"bookName"`
	Author      string     `gorm:"type:varchar(25)" json:"author"`
	PublishDate *time.Time `json:"publishDate"`
	Price       int        `gorm:"size:3" json:"price"`
}
