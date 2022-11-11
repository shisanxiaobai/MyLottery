package choose

import (
	"lottery_gateway/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoute(route *gin.RouterGroup) {
	route.POST("/do", middleware.UserJwt, Choose)
	route.GET("/result", middleware.UserJwt, ChooseResult)
}
