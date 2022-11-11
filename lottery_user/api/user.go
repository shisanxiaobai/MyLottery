package api

import (
	"context"
	"fmt"
	"lottery_user/dao"
	"lottery_user/model"
	"lottery_user/proto/user"
	"lottery_user/utils"
	"time"

	"github.com/patrickmn/go-cache"
)

type User struct {
}

var c = cache.New(60*time.Minute, 10*time.Minute)

// 注册
func (u *User) Registry(ctx context.Context, in *user.RegistryRequest, out *user.Response) error {
	code, ok := c.Get(in.Email)
	if !ok {
		out.Code = 500
		out.Msg = "未获取验证码"
		return nil
	}
	fmt.Println(code)
	fmt.Println(in.Code)
	if code != in.Code {
		out.Code = 500
		out.Msg = "验证码输入不正确，请重新输入"
		return nil
	}
	data := &model.User{
		Email:    in.Email,
		Password: in.Password,
	}
	err := dao.UserRegistry(data)
	if err != nil {
		out.Code = 500
		out.Msg = "注册用户失败"
		return err
	}
	out.Code = 200
	out.Msg = "用户注册成功"
	return nil
}

// 登录
func (u *User) Login(ctx context.Context, in *user.LoginRequest, out *user.Response) error {
	data := &model.User{
		Email:    in.Email,
		Password: in.Password,
	}
	is_ok := dao.UserLogin(data)
	if !is_ok {
		out.Code = 500
		out.Msg = "账号密码输入错误"
		return nil
	}
	out.Code = 200
	out.Msg = "登陆成功"
	return nil
}

// 发送验证码
func (u *User) SendEmail(ctx context.Context, in *user.EmailRequest, out *user.Response) error {
	email := in.Email
	is_ok := dao.UserExist(email)
	if !is_ok {
		out.Code = 500
		out.Msg = "该邮箱已被注册"
		return nil
	}
	num := utils.RandNum(6)
	c.Set(email, num, cache.DefaultExpiration)
	err := utils.SendEmail(email, num)
	if err != nil {
		out.Code = 500
		out.Msg = "发送验证码失败，请稍后再试"
		return err
	}
	out.Code = 200
	out.Msg = "发送验证码成功"
	return nil
}

// 更新积分
func (u *User) UpdateScore(ctx context.Context, in *user.UpdateRequest, out *user.Response) error {
	id := in.Id
	score := in.Score
	err := dao.UserScore(int(id), int(score))
	if err != nil {
		out.Code = 500
		out.Msg = "更新用户积分失败"
		return err
	}
	out.Code = 200
	out.Msg = "更新用户积分成功"
	return nil
}

// 查询积分
func (u *User) GetUserScore(ctx context.Context, in *user.EmailRequest, out *user.ScoreResponse) error {
	email := in.Email
	score, err := dao.GetUserScore(email)
	if err != nil {
		out.Code = 500
		out.Msg = "获取用户积分失败"
		return err
	}
	out.Code = 200
	out.Msg = "获取用户积分成功"
	out.Score = int32(score)
	return nil
}
