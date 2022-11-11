package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"type:string;not null" json:"email"`
	Password string `gorm:"type:string;not null" json:"password"`
	Score    int    `gorm:"type:int;Default:100" json:"score"`
}
