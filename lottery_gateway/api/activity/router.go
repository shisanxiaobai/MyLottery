package activity

import (
	"lottery_gateway/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoute(route *gin.RouterGroup) {
	route.GET("/list", ActiveList)
	route.POST("/add", middleware.AdminJwt, ActiveAdd)
	route.DELETE("/del", middleware.AdminJwt, ActiveDel)
	route.PUT("/edit", middleware.AdminJwt, ActiveEdit)
	route.GET("/info", ActiveInfo)
	route.GET("/prizes", ActivePrizes)
	route.POST("/addPrizes", middleware.AdminJwt, ActiveAddPrize)
}
