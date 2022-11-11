package activity

import (
	"lottery_gateway/proto/activity"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
)

var ActiveService activity.ActivityService

func init() {
	registry := consul.NewRegistry()
	service := micro.NewService(
		micro.Name("mylottery-activity"),
		micro.Registry(registry),
	)
	ActiveService = activity.NewActivityService("mylottery-activity", service.Client())
}
