package prize

import (
	"lottery_gateway/proto/prize"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
)

var PrizeService prize.PrizeService

func init() {
	registry := consul.NewRegistry()
	service := micro.NewService(
		micro.Name("mylottery-prize"),
		micro.Registry(registry),
	)
	PrizeService = prize.NewPrizeService("mylottery-prize", service.Client())
}
