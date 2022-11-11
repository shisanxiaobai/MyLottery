package model

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	Email string `gorm:"type:string;not null"`
	Name  string `gorm:"type:string;not null"`
}

type Prize struct {
	Name       string
	Num        int
	Pic        string
	CreateTime string
}
