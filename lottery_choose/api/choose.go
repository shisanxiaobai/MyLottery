package api

import (
	"context"
	"encoding/json"
	"fmt"
	"lottery_choose/dao"
	"lottery_choose/model"
	"lottery_choose/proto"
	"time"

	"github.com/streadway/amqp"
)

type Choose struct {
}

// 做活动
func (c *Choose) DoChoose(ctx context.Context, in *proto.Request, out *proto.Response) error {
	act, err := dao.GetActivityInfo(int(in.Id))
	if err != nil {
		out.Code = 500
		out.Msg = "未获取到活动信息"
		return err
	}
	now := time.Now()
	if now.Before(act.Start) || now.After(act.End) {
		out.Code = 500
		out.Msg = "不在活动时间之内"
		return err
	}
	if !dao.Check(in.Email, act.Cost) {
		out.Code = 500
		out.Msg = "积分不足"
		return nil
	}
	prize, err := dao.Choose(int(in.Id))
	if err != nil {
		dao.Recover(in.Email, act.Cost)
		out.Code = 500
		out.Msg = "奖品已抽完"
		return err
	}
	data := &model.Order{
		Email: in.Email,
		Name:  prize,
	}
	err = dao.OrderAdd(data)
	if err != nil {
		out.Code = 500
		out.Msg = "创建订单失败"
		return err
	}
	out.Code = 200
	out.Msg = "抽奖成功"
	out.Name = prize
	return nil
}

func init() {
	orders, err := dao.Ch.Consume("order_queue", "order_consumer", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	for order := range orders {
		go OrderApply(order)
	}
}

func OrderApply(order amqp.Delivery) {
	body := string(order.Body)
	byte_data := map[string]interface{}{}
	err := json.Unmarshal([]byte(body), &byte_data)
	if err != nil {
		panic(err)
	}
	email := byte_data["email"].(string)
	id := int(byte_data["id"].(float64))
	act, err := dao.GetActivityInfo(id)
	if err != nil {
		fmt.Println(err)
		order.Ack(true)
		rep_map := map[string]interface{}{
			"code": 500,
			"msg":  "未获取到活动信息",
		}
		rep_data, _ := json.Marshal(rep_map)
		dao.Set(email, string(rep_data))
		return
	}
	now := time.Now()
	if now.Before(act.Start) || now.After(act.End) {
		fmt.Println(now)
		fmt.Println(act.Start)
		fmt.Println(act.End)
		order.Ack(true)
		rep_map := map[string]interface{}{
			"code": 500,
			"msg":  "不在活动时间内",
		}
		rep_data, _ := json.Marshal(rep_map)
		dao.Set(email, string(rep_data))
		return
	}
	if !dao.Check(email, act.Cost) {
		order.Ack(true)
		rep_map := map[string]interface{}{
			"code": 500,
			"msg":  "积分不足",
		}
		rep_data, _ := json.Marshal(rep_map)
		dao.Set(email, string(rep_data))
		return
	}
	prize, err := dao.Choose(id)
	fmt.Println(prize)
	if err != nil {
		order.Ack(true)
		rep_map := map[string]interface{}{
			"code": 500,
			"msg":  "奖品已抽完",
		}
		dao.Recover(email, act.Cost)
		rep_data, _ := json.Marshal(rep_map)
		dao.Set(email, string(rep_data))
		return
	}
	data := &model.Order{
		Email: email,
		Name:  prize,
	}
	err = dao.OrderAdd(data)
	if err != nil {
		order.Ack(true)
		rep_map := map[string]interface{}{
			"code": 500,
			"msg":  "创建订单失败",
		}
		rep_data, _ := json.Marshal(rep_map)
		dao.Set(email, string(rep_data))
		return
	}
	order.Ack(true)
	rep_map := map[string]interface{}{
		"code":  200,
		"msg":   "成功",
		"prize": prize,
	}
	rep_data, _ := json.Marshal(rep_map)
	dao.Set(email, string(rep_data))
	return
}
