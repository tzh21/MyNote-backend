package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	// 注册
	userHandler := NewUserHandler()
	router.POST("/signup", userHandler.HandleSignup)

	// 登录
	router.POST("/login", userHandler.HandleLogin)

	// 仅管理员可以访问的 api
	adminGroup := router.Group("/admin", auth())
	{
		// 获取所有用户的信息
		adminGroup.GET("/user", userHandler.HandleGetAllUsers)
	}
}

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO
		c.Next()
	}
}
