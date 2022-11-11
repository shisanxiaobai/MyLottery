package main

import (
	"log"
	"lottery_user/api"

	"lottery_user/proto/admin"
	"lottery_user/proto/user"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/consul/v2"

	"github.com/micro/go-micro/v2/registry"
)

func main() {
	//consul注册件
	newRegistry := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)
	//micro微服务实例
	service := micro.NewService(
		micro.Name("mylottery-user"),
		micro.Version("latest"),
		micro.Registry(newRegistry),
	)

	//初始化服务
	service.Init()

	//注册user和admin的处理函数
	user.RegisterUserHandler(service.Server(), new(api.User))
	admin.RegisterAdminHandler(service.Server(), new(api.Admin))
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
