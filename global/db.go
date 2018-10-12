package global

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDB()  {
	db, _ := gorm.Open("mysql", "root:mysql.com2017@tcp(106.14.96.72:3306)/common_db?charset=utf8&parseTime=True&loc=Local")
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	DB=db
}
