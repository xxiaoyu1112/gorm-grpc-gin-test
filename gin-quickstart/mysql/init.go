package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const dsn_mysql = "root:123456@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"

var (
	DB *gorm.DB
)

func init() {
	db, err := gorm.Open("mysql", dsn_mysql)
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
}
