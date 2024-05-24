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
	router.POST("/avatar-name/:username", profileHandler.HandleUploadAvatarPath)
	router.GET("/avatar-name/:username", profileHandler.HandleGetAvatarPath)
	router.POST("/avatar/:username/:filename", profileHandler.HandleUploadAvatarFile)
	router.GET("/avatar/:username/:filename", profileHandler.HandleGetAvatarFile)
	// 全部用户信息
	router.GET("/profile/:username", profileHandler.HandleGetProfile)

	// 文件相关路由
	fileHandler := NewFileHandler(db)
	router.GET("/list/:user", fileHandler.HandleGetList)
	router.POST("/blocks/:username/:filename", fileHandler.HandleUploadBlocks)
	router.GET("/blocks/:username/:filename", fileHandler.HandleGetBlocks)
	router.POST("/image/:username/:filename", fileHandler.HandleUploadImage)
	router.GET("/image/:username/:filename", fileHandler.HandleGetImage)
	router.POST("/audio/:username/:filename", fileHandler.HandleUploadAudio)
	router.GET("/audio/:username/:filename", fileHandler.HandleGetAudio)

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
