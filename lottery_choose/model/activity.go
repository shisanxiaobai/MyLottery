package model

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	Cost  int `gorm:"type:int;Default:100"`
	Start time.Time
	End   time.Time
}
