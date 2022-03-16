package mysql

func Close() {
	DB.Close()
}

// 模型绑定
//DB.AutoMigrate(&Todo{}) //  表名为todos
