package choose

import (
	"lottery_gateway/proto/choose"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
)

var ChooseService choose.ChooseService

func init() {
	registry := consul.NewRegistry()
	service := micro.NewService(
		micro.Name("mylottery-choose"),
		micro.Registry(registry),
	)
	ChooseService = choose.NewChooseService("mylottery-choose", service.Client())
}
