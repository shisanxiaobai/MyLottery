package prize

import (
	"lottery_gateway/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoute(route *gin.RouterGroup) {
	route.GET("/list", PrizeList)
	route.POST("/add", middleware.AdminJwt, PrizeAdd)
	route.DELETE("/del", middleware.AdminJwt, PrizeDel)
	route.PUT("/edit", middleware.AdminJwt, PrizeEdit)
	route.GET("/info", PrizeInfo)
}
