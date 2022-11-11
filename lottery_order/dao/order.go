package dao

import (
	"lottery_order/model"
	"strconv"
)

// 获得订单列表
func OrderList(pageNum int, pageSize int) (orders []*model.Order, count int64, err error) {
	err = Db.Model(&model.Order{}).Count(&count).Error
	err = Db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Order("created_at Desc").Find(&orders).Error
	return
}

// 根据email查询订单
func OrderByEmail(email string, pageNum int, pageSize int) (orders []*model.Order, count int64, err error) {
	err = Db.Model(&model.Order{}).Where("Email=?", email).Count(&count).Error
	err = Db.Where("Email=?", email).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&orders).Error
	return
}

// 添加订单
func OrderAdd(order *model.Order) error {
	return Db.Create(&order).Error
}

func OrderDel(id int) error {
	return Db.Delete(&model.Order{}, id).Error
}

func PrizeInfo(name string) (prize *model.Prize, err error) {
	value, _ := Rdb.HGetAll(name).Result()
	num, _ := strconv.Atoi(value["num"])
	prize = &model.Prize{
		Name:       name,
		Num:        num,
		Pic:        value["pic"],
		CreateTime: value["createTime"],
	}
	return
}
