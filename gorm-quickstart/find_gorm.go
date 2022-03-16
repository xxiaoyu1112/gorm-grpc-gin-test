package main

import "fmt"

type UserInfo struct {
	Name string
	Age string
}

func Find() {
	// 方式一：map查找
	//result := map[string]interface{}{}
	//GLOBAL_DB.Model(&TestUser{}).First(&result)
	//fmt.Println(result)

	//// 方式二：结构体查找
	//var User TestUser
	//GLOBAL_DB.Model(&TestUser{}).Take(&User)
	//fmt.Println(User)
	// where

	//var User []TestUser

	// 字符串的形式
	//GLOBAL_DB.Where("name = ? AND age = ?","xing",21).
	//	Or("name = ?","ding").First(&User)

	//  map形式
	//GLOBAL_DB.Where(map[string]interface{}{
	//	"name":"ning",
	//	"age":19,
	//}).First(&User)

	// 结构体形式
	//GLOBAL_DB.Where(TestUser{Name: "ding"}).First(&User)

	//GLOBAL_DB.First(&User,map[string]interface{}{
	//		"name":"ning",
	//		"age":19,
	//	})
	//GLOBAL_DB.Omit("name").Where("name LIKE ?","%in%").Find(&User)

	var u []UserInfo
	GLOBAL_DB.Model(&TestUser{}).Where("name LIKE ?","%in%").Find(&u)
	fmt.Println(u)

}
