package dao

import "lottery_choose/model"

func OrderAdd(order *model.Order) error {
	return Db.Create(&order).Error
}
