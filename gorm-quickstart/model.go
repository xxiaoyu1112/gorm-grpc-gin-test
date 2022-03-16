package main

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	UUID  uint  `gorm:"primaryKey"`
	Time  time.Time  `gorm:"column:my_time"`
}

type TestUser struct {
	gorm.Model
	Name         string    `gorm:"default:ming"`
	Age          uint8		`gorm:"comment:年龄"`
}

func UserCreate() {
	GLOBAL_DB.AutoMigrate(&TestUser{})
}
