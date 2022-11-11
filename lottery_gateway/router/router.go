package router

import (
	"lottery_gateway/api/activity"
	"lottery_gateway/api/choose"
	"lottery_gateway/api/order"
	"lottery_gateway/api/prize"
	"lottery_gateway/api/user"
	"lottery_gateway/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	router.Use(middleware.Cors())
	router.Use(gin.Recovery())
	router.Use(middleware.Logger())
	router_user := router.Group("/user")
	router_prize := router.Group("/prize")
	router_order := router.Group("/order")
	router_choose := router.Group("/choose")
	router_activity := router.Group("/activity")

	user.InitRoute(router_user)
	prize.InitRoute(router_prize)
	order.InitRoute(router_order)
	choose.InitRoute(router_choose)
	activity.InitRoute(router_activity)

	router.Static("/upload", "./upload")
}
