package dao

import "lottery_choose/model"

func GetActivityInfo(id int) (act *model.Activity, err error) {
	err = Db.First(&act, id).Error
	return
}
