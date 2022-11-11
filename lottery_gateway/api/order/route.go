package order

import (
	"lottery_gateway/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoute(route *gin.RouterGroup) {
	route.GET("/list", middleware.UserJwt, OrderList)
	route.GET("/userList", middleware.UserJwt, OrderByUser)
	route.DELETE("/del", middleware.UserJwt, OrderDel)
}
