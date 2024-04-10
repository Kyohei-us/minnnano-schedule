package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id   int    `gorm:"" json:"id"`
	Name string `gorm:"" json:"name"`
}
