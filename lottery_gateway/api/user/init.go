package user

import (
	"lottery_gateway/proto/admin"
	"lottery_gateway/proto/user"
	"lottery_gateway/wrapper"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
)

var (
	UserService  user.UserService
	AdminService admin.AdminService
)

func init() {
	registry := consul.NewRegistry()
	service := micro.NewService(
		micro.Name("mylottery-user"),
		micro.Registry(registry),
		micro.WrapClient(wrapper.NewUserWrapper),
	)
	UserService = user.NewUserService("mylottery-user", service.Client())
	AdminService = admin.NewAdminService("mylottery-user", service.Client())
}
