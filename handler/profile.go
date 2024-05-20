package handler

import (
	"MyNote-backend/model"
	"bytes"
	"io"
	"mime/multipart"
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

func (h *ProfileHandler) HandleGetProfile(c *gin.Context) {
	username := c.Param("username")

	var profile model.Profile
	if err := h.db.Where("username = ?", username).First(&profile).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	// 创建一个 bytes.Buffer 用于构建 multipart form 数据
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	// 写入 motto 和 nickname 字段
	w.WriteField("motto", profile.Motto)
	w.WriteField("nickname", profile.Nickname)

	// 写入 avatar 图片文件
	// 假设 avatar 数据存储在 profile.Avatar 字段中
	if profile.Avatar != nil {
		// 创建文件部分
		fw, err := w.CreateFormFile("avatar", "avatar")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create form file"})
			return
		}
		// 将图片数据写入文件部分
		if _, err := fw.Write(profile.Avatar); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write avatar data"})
			return
		}
	}

	// 关闭 multipart writer
	w.Close()

	// 设置 Content-Type 为 multipart/form-data
	c.Writer.Header().Set("Content-Type", w.FormDataContentType())

	// 将 multipart form 数据写入响应主体
	c.Writer.Write(b.Bytes())
}

func (h *ProfileHandler) HandleUploadAvatar(c *gin.Context) {
	err := c.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	username := c.Param("username")

	file, _, err := c.Request.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	// 读取上传的文件内容为字节切片
	avatarData, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 更新数据库中的 Avatar 字段
	var profile model.Profile
	if err := h.db.Where("username = ?", username).First(&profile).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}
	profile.Avatar = avatarData
	if err := h.db.Save(&profile).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Avatar uploaded successfully"})
}

func (h *ProfileHandler) HandleGetAvatar(c *gin.Context) {
	username := c.Param("username")

	// 查询数据库，获取用户的 Profile 记录
	var profile model.Profile
	if err := h.db.Where("username = ?", username).First(&profile).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	// 检查用户是否有头像数据
	if len(profile.Avatar) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Avatar not found"})
		return
	}

	// 将头像数据直接返回给客户端
	c.Data(http.StatusOK, "image/jpeg", profile.Avatar)
}
