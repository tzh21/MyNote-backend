package handler

import (
	"github.com/gin-gonic/gin"
)

type LoginHandler struct{}

func NewLoginHandler() *LoginHandler {
	return &LoginHandler{}
}

func (h *LoginHandler) LoginHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "login",
	})
}
