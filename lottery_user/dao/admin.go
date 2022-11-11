package dao

import "lottery_user/model"

func AdminLogin(data *model.Admin) bool {
	var count int64
	err := Db.Model(model.Admin{}).Where("Username=? AND Password=?", data.Username, data.Password).Count(&count).Error
	return err == nil && count > 0
}
