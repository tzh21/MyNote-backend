package handler

import (
	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) HandleSignup(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "signup",
	})
}

func (h *UserHandler) HandleLogin(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "login",
	})
}

func (h *UserHandler) HandleGetAllUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "user",
	})
}
