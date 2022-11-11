package api

import (
	"context"
	"lottery_user/dao"
	"lottery_user/model"
	"lottery_user/proto/admin"
)

type Admin struct {
}

// 管理员登录
func (a *Admin) AdminLogin(ctx context.Context, in *admin.LoginRequest, out *admin.LoginResponse) error {
	data := &model.Admin{
		Username: in.Username,
		Password: in.Password,
	}
	is_ok := dao.AdminLogin(data)
	if !is_ok {
		out.Code = 500
		out.Msg = "账号密码输入错误"
		return nil
	}
	out.Code = 200
	out.Msg = "登陆成功"
	return nil
}

// 用户列表
func (a *Admin) UserList(ctx context.Context, in *admin.UsersRequest, out *admin.UsersResponse) error {
	count, users, err := dao.UserList(int(in.PageNum), int(in.PageSize))
	if err != nil {
		out.Code = 500
		out.Msg = "用户列表获取失败"
		return err
	}
	rep_datas := []*admin.UserInfo{}
	for _, v := range users {
		rep_data := &admin.UserInfo{
			Id:         int32(v.ID),
			Email:      v.Email,
			Score:      int32(v.Score),
			CreateTime: v.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		rep_datas = append(rep_datas, rep_data)
	}
	out.Users = rep_datas
	out.Code = 200
	out.Msg = "用户列表获取成功"
	out.Count = int32(count)
	return nil
}

// 删除用户
func (a *Admin) UserDel(ctx context.Context, in *admin.DelRequest, out *admin.DelResponse) error {
	err := dao.UserDel(int(in.Id))
	if err != nil {
		out.Code = 500
		out.Msg = "用户删除失败"
		return err
	}
	out.Code = 200
	out.Msg = "用户删除成功"
	return nil
}
