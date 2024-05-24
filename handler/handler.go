package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	// 登录相关路由
	userHandler := NewUserHandler(db)
	router.POST("/signup", userHandler.HandleSignup)
	router.POST("/login", userHandler.HandleLogin)

	// 用户信息相关路由
	profileHandler := NewProfileHandler(db)
	// 个性签名
	router.POST("/motto/:username", profileHandler.HandleUploadMotto)
	router.GET("/motto/:username", profileHandler.HandleGetMotto)
	// 昵称
	router.POST("/nickname/:username", profileHandler.HandleUploadNickname)
	router.GET("/nickname/:username", profileHandler.HandleGetNickname)
	// 头像
	router.POST("/avatar/:username", profileHandler.HandleUploadAvatar)
	router.GET("/avatar/:username", profileHandler.HandleGetAvatar)
	// 全部用户信息
	router.GET("/profile/:username", profileHandler.HandleGetProfile)

	// 文件相关路由
	fileHandler := NewFileHandler(db)
	router.GET("/list/:user", fileHandler.HandleGetList) // 返回用户所有文件路径组成的列表。每个元素的结构为 username/category/filename。客户端接收后根据表的内容逐个请求文件。
	// path 根目录为用户名，比如笔记的结构为 username/category/filename
	// router.POST("/upload/:username/:filename", fileHandler.HandleUpload)
	router.POST("/blocks/:username/:filename", fileHandler.HandleUploadBlocks)
	router.GET("/blocks/:username/:filename", fileHandler.HandleGetBlocks)
	router.POST("/image/:username/:filename", fileHandler.HandleUploadImage)
	router.GET("/image/:username/:filename", fileHandler.HandleGetImage)
	// router.GET("/download/*path", fileHandler.HandleDownload)

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
