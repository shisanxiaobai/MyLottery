package model

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	Name  string `gorm:"type:string;not null"`
	Desc  string `gorm:"type:string;not null"`
	Pic   string `gorm:"type:string;not null"`
	User  string `gorm:"type:string;not null"`
	Cost  int    `gorm:"type:int;Default:100"`
	Type  int    `gorm:"type:int;Default:0"`
	Start time.Time
	End   time.Time
}
