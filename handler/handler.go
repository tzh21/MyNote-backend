package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	// 用户相关路由
	userHandler := NewUserHandler(db)
	router.POST("/signup", userHandler.HandleSignup)
	router.POST("/login", userHandler.HandleLogin)

	// 文件相关路由
	fileHandler := NewFileHandler(db)
	router.GET("/list/:user", fileHandler.HandleList) // 返回用户所有文件路径组成的列表。每个元素的结构为 username/category/filename。客户端接收后根据表的内容逐个请求文件。
	// path 根目录为用户名，比如笔记的结构为 username/category/filename
	router.POST("/upload/*path", fileHandler.HandleUpload)
	router.GET("/download/*path", fileHandler.HandleDownload)

	// 仅管理员可以访问的 api
	adminGroup := router.Group("/admin", auth())
	{
		// 获取所有用户的信息
		adminGroup.GET("/user", userHandler.HandleGetAllUsers)
	}
}

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO 完善 auth 中间件
		c.Next()
	}
}
