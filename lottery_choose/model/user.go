package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"type:string;not null"`
	Password string `gorm:"type:string;not null"`
	Score    int    `gorm:"type:int;Default:100"`
}
