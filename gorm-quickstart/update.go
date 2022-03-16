package main

import "fmt"

func Update() {
	//  update   只更新选择的字段
	//  updates   更新所有字段  此时有两种形式  一种为map  一种为结构体  结构体零值不参与更新
	//  save      无论如何都更新  所有的内容  包括 零值

	// update
	//GLOBAL_DB.Model(&TestUser{}).Where(
	//	"name LIKE ?","%in%").Update("name","inn")

	// save
	//var users []TestUser
	//dbRes := GLOBAL_DB.Where("name = ?","inn").Find(&users)
	//for k := range users{
	//	users[k].Age = 18
	//}
	//dbRes.Save(&users)

	// Updates
	// 结构体
	var user TestUser
	//GLOBAL_DB.Find(&user).Updates(TestUser{Name: "ming",Age: 0})
	//fmt.Println(&user)

	// map  可以更新零值
	GLOBAL_DB.Find(&user).Updates(map[string]interface{}{"Name": "ming","Age": 0})
	fmt.Println(&user)
}
