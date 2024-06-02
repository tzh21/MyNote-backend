package handler

import (
	"MyNote-backend/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FileHandler struct {
	db *gorm.DB
}

func NewFileHandler(db *gorm.DB) *FileHandler {
	return &FileHandler{db: db}
}

func (h *FileHandler) HandleUploadBlocks(c *gin.Context) {
	username := c.Param("username")
	filename := c.Param("filename")
	savePath := utils.BlocksPath(username, filename)
	utils.HandleUpload(c, savePath)
}

func (h *FileHandler) HandleGetBlocks(c *gin.Context) {
	username := c.Param("username")
	filename := c.Param("filename")
	path := utils.BlocksPath(username, filename)

	c.File(path)
}

func (h *FileHandler) HandleUploadImage(c *gin.Context) {
	username := c.Param("username")
	filename := c.Param("filename")
	savePath := utils.ImagesPath(username, filename)
	utils.HandleUpload(c, savePath)
}

func (h *FileHandler) HandleGetImage(c *gin.Context) {
	username := c.Param("username")
	filename := c.Param("filename")
	path := utils.ImagesPath(username, filename)

	c.File(path)
}

func (h *FileHandler) HandleUploadAudio(c *gin.Context) {
	username := c.Param("username")
	filename := c.Param("filename")
	savePath := utils.AudioPath(username, filename)
	utils.HandleUpload(c, savePath)
}

func (h *FileHandler) HandleGetAudio(c *gin.Context) {
	username := c.Param("username")
	filename := c.Param("filename")
	path := utils.AudioPath(username, filename)

	c.File(path)
}

func getNoteFileNames(username string) ([]string, error) {
	var dir = utils.BlocksBasePath(username)

	var fileNames []string
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range entries {
		if !file.IsDir() {
			fileNames = append(fileNames, file.Name())
		}
	}

	return fileNames, nil
}

func (h *FileHandler) HandleGetList(c *gin.Context) {
	username := c.Param("user")

	fileNames, err := getNoteFileNames(username)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"files": fileNames})
}

// 删除 blocks 和依赖文件
func (h *FileHandler) HandleDeleteNote(c *gin.Context) {
	username := c.Param("username")
	filename := c.Param("filename")

	blocksPath := utils.BlocksPath(username, filename)

	if err := os.Remove(blocksPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}
