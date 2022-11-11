package main

import (
	"log"
	"lottery_activity/api"
	"lottery_activity/proto"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	newRegistry := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)
	service := micro.NewService(
		micro.Name("mylottery-activity"),
		micro.Version("latest"),
		micro.Registry(newRegistry),
	)
	proto.RegisterActivityHandler(service.Server(), new(api.Activity))
	service.Init()
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
