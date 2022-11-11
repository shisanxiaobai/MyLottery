package dao

import "lottery_activity/model"

func ActiveAdd(active *model.Activity) (int, error) {
	return int(active.ID), Db.Create(&active).Error
}

func ActiveDel(id int) error {
	err := Db.Delete(&model.Activity{}, id).Error
	if err != nil {
		return err
	}
	return ActiveDelPrizes(id)
}

func ActiveEdit(active *model.Activity, id int) error {
	return Db.Model(&model.Activity{}).Where("ID=?", id).Updates(&active).Error
}

func ActiveList(pageNum, pageSize int) (actives []*model.Activity, count int64, err error) {
	Db.Model(&model.Activity{}).Count(&count)
	err = Db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&actives).Error
	return
}

func ActiveInfo(id int) (active *model.Activity, err error) {
	err = Db.Where("ID=?", id).First(&active).Error
	return
}
