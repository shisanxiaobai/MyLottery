package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Username string `gorm:"type:string;not null" json:"username"`
	Password string `gorm:"type:string;not null" json:"password"`
}
