package handler

import (
	"MyNote-backend/model"
	"MyNote-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProfileHandler struct {
	db *gorm.DB
}

func NewProfileHandler(db *gorm.DB) *ProfileHandler {
	return &ProfileHandler{db: db}
}

// 用户信息包括：个性签名、头像、头像地址
// 要求用户信息具有可扩展性

func (h *ProfileHandler) HandleUploadMotto(c *gin.Context) {
	username := c.Param("username")

	// 从请求中解析 motto
	type request struct {
		Motto string `json:"motto"`
	}
	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 更新 motto
	var profile model.Profile
	if err := h.db.Where("username = ?", username).FirstOrCreate(&profile,
		model.Profile{
			Username: username,
		}).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	profile.Motto = req.Motto
	if err := h.db.Save(&profile).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Motto updated successfully"})
}

func (h *ProfileHandler) HandleGetMotto(c *gin.Context) {
	username := c.Param("username")

	var profile model.Profile
	if err := h.db.Where("username = ?", username).First(&profile).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"motto": profile.Motto})
}

func (h *ProfileHandler) HandleUploadNickname(c *gin.Context) {
	username := c.Param("username")

	// 从请求中解析 nickname
	type request struct {
		Nickname string `json:"nickname"`
	}
	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 更新 nickname
	var profile model.Profile
	if err := h.db.Where("username = ?", username).FirstOrCreate(&profile,
		model.Profile{
			Username: username,
		}).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	profile.Nickname = req.Nickname
	if err := h.db.Save(&profile).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Nickname updated successfully"})
}

func (h *ProfileHandler) HandleGetNickname(c *gin.Context) {
	username := c.Param("username")

	var profile model.Profile
	if err := h.db.Where("username = ?", username).First(&profile).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"nickname": profile.Nickname})
}

func (h *ProfileHandler) HandleUploadAvatarPath(c *gin.Context) {
	username := c.Param("username")

	// 从请求中解析 avatar
	type request struct {
		Avatar string `json:"avatar"`
	}
	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 更新 avatar
	var profile model.Profile
	if err := h.db.Where("username = ?", username).FirstOrCreate(&profile,
		model.Profile{
			Username: username,
		}).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	profile.Avatar = req.Avatar
	if err := h.db.Save(&profile).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Avatar updated successfully"})
}

func (h *ProfileHandler) HandleGetAvatarPath(c *gin.Context) {
	username := c.Param("username")

	var profile model.Profile
	if err := h.db.Where("username = ?", username).First(&profile).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"avatar": profile.Avatar})
}

func (h *ProfileHandler) HandleUploadAvatarFile(c *gin.Context) {
	username := c.Param("username")
	filename := c.Param("filename")
	savePath := utils.AvatarPath(username, filename)
	utils.HandleUpload(c, savePath)
}

func (h *ProfileHandler) HandleGetAvatarFile(c *gin.Context) {
	username := c.Param("username")
	filename := c.Param("filename")
	path := utils.AvatarPath(username, filename)

	c.File(path)
}

func (h *ProfileHandler) HandleGetProfile(c *gin.Context) {
	username := c.Param("username")

	var profile model.Profile
	if err := h.db.Where("username = ?", username).First(&profile).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"username": profile.Username,
		"motto":    profile.Motto,
		"nickname": profile.Nickname,
		"avatar":   profile.Avatar,
	})
}

func (h *ProfileHandler) HandleChangePassword(c *gin.Context) {
	username := c.Param("username")

	// 从请求中解析 password
	type request struct {
		Password string `json:"password"`
	}
	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 更新 password
	var user model.User
	if err := h.db.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	user.Password = req.Password
	if err := h.db.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}
