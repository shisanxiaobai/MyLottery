package dao

import (
	"fmt"
	"lottery_prize/model"
	"strconv"
	"time"
)

func getRedisName(key string) string {
	return fmt.Sprintf("lottery:prize:%s", key)
}

// 奖品列表
func PrizeList() (prizes []*model.Prize, err error) {
	list_name, err := Rdb.SMembers(getRedisName("")).Result()
	for _, v := range list_name {
		value, _ := Rdb.HGetAll(v).Result()
		num, _ := strconv.Atoi(value["num"])
		prize := &model.Prize{
			Name:       v,
			Num:        num,
			Pic:        value["pic"],
			CreateTime: value["createTime"],
		}
		prizes = append(prizes, prize)
	}
	return
}

// 添加奖品
func PrizeAdd(prize *model.Prize) error {
	name := prize.Name
	err := Rdb.SAdd(getRedisName(""), name).Err()
	if err != nil {
		fmt.Println(err)
		return err
	}
	if err != nil {
		fmt.Println(err)
		return err
	}
	Rdb.HSet(name, "num", prize.Num)
	Rdb.HSet(name, "pic", prize.Pic)
	Rdb.HSet(name, "createTime", prize.CreateTime)
	return nil
}

// 删除奖品
func PrizeDel(name string) error {
	err := Rdb.SRem(getRedisName(""), name).Err()
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	Rdb.HDel(name, "num")
	Rdb.HDel(name, "pic")
	Rdb.HDel(name, "createTime")
	return nil
}

// 编辑奖品
func PrizeEdit(name string, prize *model.Prize) error {
	err := Rdb.SRem(getRedisName(""), name).Err()
	if err != nil {
		return err
	}
	Rdb.HDel(name, "num")
	Rdb.HDel(name, "pic")
	createTime, err := Rdb.HGet(name, "createTime").Result()
	if err != nil {
		createTime = time.Now().Format("2006-01-02 15:04:05")
	}
	Rdb.HDel(name, "createTime")
	err = Rdb.SAdd(getRedisName(""), prize.Name).Err()
	if err != nil {
		return err
	}
	Rdb.HSet(prize.Name, "num", prize.Num)
	if err != nil {
		return err
	}
	Rdb.HSet(prize.Name, "pic", prize.Pic)
	Rdb.HSet(prize.Name, "createTime", createTime)
	return nil
}

// 奖品信息
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
