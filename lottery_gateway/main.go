package main

import (
	"log"
	"lottery_gateway/router"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	r := gin.New()
	router.InitRouter(r)
	newRegistry := consul.NewRegistry(
		registry.Addrs("127.0.0.1"),
	)
	service := web.NewService(
		web.Name("mylottery-gateway"),
		web.Version("latest"),
		web.Handler(r),
		web.Registry(newRegistry),
		web.Address("127.0.0.1:9002"),
	)
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
