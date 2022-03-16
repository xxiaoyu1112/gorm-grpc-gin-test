package main

import "fmt"

func Create() {
	dbres := GLOBAL_DB.Create(&[]TestUser{
		{Name: "ming", Age: 18,},
		{Name: "ning", Age: 19,},
		{Name: "ding", Age: 20,},
		{Name: "xing", Age: 21,},
	})
	fmt.Println(dbres.Error,dbres.RowsAffected)
	if dbres.Error!= nil {
		fmt.Println("创建失败")
	}else {
		fmt.Println("创建成功")
	}
}
