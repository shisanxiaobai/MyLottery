package dao

import (
	"encoding/json"
	"fmt"
	"lottery_activity/model"
	"strconv"
)

func getRedisName(index int) string {
	key := strconv.Itoa(index)
	return fmt.Sprintf("lottery:prize:%v", key)
}

func ActivePrizes(id int) (prizes []*model.Prize, err error) {
	list_name, err := Rdb.SMembers(getRedisName(id)).Result()
	fmt.Println(list_name)
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

func ActiveAddPrize(id int, names string) error {
	values := []string{}
	err := json.Unmarshal([]byte(names), &values)
	if err != nil {
		return err
	}
	for _, v := range values {
		if err := Rdb.SAdd(getRedisName(id), v).Err(); err != nil {
			return err
		}
	}
	return nil
}

func ActiveDelPrizes(id int) error {
	list_name, err := Rdb.SMembers(getRedisName(id)).Result()
	if err != nil {
		return err
	}
	for _, v := range list_name {
		err := ActiveDelPrize(id, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func ActiveDelPrize(id int, name string) error {
	return Rdb.SRem(getRedisName(id), name).Err()
}
