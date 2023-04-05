package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/morningcat2018/book-service/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

func init() {
	db, _ = gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:@tcp(127.0.0.1:3306)/db_go_test?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         256,                                                                          // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                         // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                         // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                         // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                        // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})

	db.AutoMigrate(&entity.Book{})
	println("db init over")
}

func BookInsert(book *entity.Book) {
	//book.LogicStatus = 1
	db.Save(book)
}

func BookGetById(id string) entity.Book {
	// Read
	var book entity.Book
	db.First(&book, id) // 根据整型主键查找
	//db.First(&book, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
	if book.IsEmpty() {
		return book
	}
	//if book.LogicStatus == 0 {
	//	return entity.Book{}
	//}
	if book.DeletedAt.Valid {
		return entity.Book{}
	}
	return book
}

func BookUpdateById(id string) entity.Book {
	// Read
	var book entity.Book
	db.First(&book, id) // 根据整型主键查找
	if book.IsEmpty() {
		return book
	}

	// Update - 将 book 的 LogicStatus 更新为 0
	//db.Model(&book).Update("LogicStatus", 0)
	book.DeletedAt.Valid = true
	book.DeletedAt.Time = time.Now()
	db.Updates(book)

	return book
}
