package dao

import (
	"lottery_choose/model"

	"gorm.io/gorm"
)

func Check(email string, cost int) bool {
	var user model.User
	err := Db.Select("Score").Where("Email=?", email).First(&user).Error
	if err != nil {
		return false
	}
	if user.Score < cost {
		return false
	}
	err = Db.Model(&model.User{}).Where("Email=?", email).UpdateColumn("Score", gorm.Expr("Score-?", cost)).Error
	return err == nil
}

func Recover(email string, cost int) {
	Db.Model(&model.User{}).Where("Email=?", email).UpdateColumn("Score", gorm.Expr("Score+?", cost))
}
