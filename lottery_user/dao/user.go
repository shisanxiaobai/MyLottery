package dao

import "lottery_user/model"

func UserRegistry(data *model.User) error {
	return Db.Create(&data).Error
}

func UserLogin(data *model.User) bool {
	var count int64
	err := Db.Model(&model.User{}).Where("Email=? AND Password=?", data.Email, data.Password).Count(&count).Error
	return err == nil && count > 0
}

func UserExist(email string) bool {
	var count int64
	err := Db.Where("Email=?", email).Find(&model.User{}).Count(&count).Error
	return count == 0 && err == nil
}

func UserList(pageNum int, pageSize int) (count int64, users []*model.User, err error) {
	Db.Model(model.User{}).Count(&count)
	err = Db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&users).Error
	return
}

func UserDel(id int) error {
	return Db.Delete(&model.User{}, id).Error
}

func UserScore(id int, score int) error {
	return Db.Model(model.User{}).Where("ID=?", id).Update("Score", score).Error
}

func GetUserScore(email string) (int, error) {
	var user model.User
	err := Db.Where("Email=?", email).First(&user).Error
	return user.Score, err
}
