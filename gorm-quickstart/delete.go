package main

import "fmt"

func Delete() {
	var user TestUser

	// 软删除 不删除字段， 而是给 DeleteAt字段 添加时间
	//GLOBAL_DB.Where("name = ? ","ming").Delete(&user)

	// 直接删除
	//GLOBAL_DB.Unscoped().Where("name = ? ","ming").Delete(&user)

	GLOBAL_DB.Raw("Select * FRom testusers where name = ? ","ming").Scan(&user)
	fmt.Println(&user)
}
