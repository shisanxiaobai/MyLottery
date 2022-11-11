package main

import (
	"log"
	"lottery_choose/proto"

	"lottery_choose/api"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"

	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	newRegistry := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)
	service := micro.NewService(
		micro.Name("mylottery-choose"),
		micro.Version("latest"),
		micro.Registry(newRegistry),
	)
	service.Init()
	proto.RegisterChooseHandler(service.Server(), new(api.Choose))
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
