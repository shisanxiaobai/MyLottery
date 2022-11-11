package order

import (
	"lottery_gateway/proto/order"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
)

var OrderService order.OrderService

func init() {
	registry := consul.NewRegistry()
	service := micro.NewService(
		micro.Name("mylottery-order"),
		micro.Registry(registry),
	)
	OrderService = order.NewOrderService("mylottery-order", service.Client())
}
