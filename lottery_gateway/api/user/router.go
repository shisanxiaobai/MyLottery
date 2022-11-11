package user

import (
	"lottery_gateway/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoute(router *gin.RouterGroup) {
	router.POST("registry", Registry)
	router.POST("login", Login)
	router.POST("email", SendEmail)
	router.GET("score", middleware.UserJwt, GetScore)

	router.POST("adminLogin", AdminLogin)
	router.GET("list", middleware.AdminJwt, UserList)
	router.DELETE("del", middleware.AdminJwt, UserDel)
	router.PUT("edit", middleware.AdminJwt, UpdateScore)
}
