package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strings"
	"time"
)

var GLOBAL_DB *gorm.DB

func main() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:123456@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize: 171, // string 类型字段的默认长度
		DisableDatetimePrecision: true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex: true,
		DontSupportRenameColumn: true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置

	}), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix: "gva_",   // 表名前缀，‘User’的表名应该是‘t_users’
			//SingularTable: false, // 使用单数表名，启用该选项，此时，‘User’ 的表名应该会‘t_user’，不启用则为‘t_users’
			NoLowerCase: true, // skip the snake_casing of names
			NameReplacer: strings.NewReplacer("CID", "Cid"),
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 逻辑外检（代码里面自动提现外键关系）
	})

	fmt.Println(db,err)

	sqlDB,_ := db.DB()
	sqlDB.SetMaxIdleConns(10)   // 连接池中最大的空闲连接数
	sqlDB.SetMaxOpenConns(100)  //  连接池最多可容纳的连接数量
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接池中连接的最大可复用时间

	GLOBAL_DB = db

	UserCreate()
	//Create()
	//Find()
	//Update()
	//Delete()

}
