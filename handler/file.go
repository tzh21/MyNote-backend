package handler

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FileHandler struct {
	db *gorm.DB
}

func NewFileHandler(db *gorm.DB) *FileHandler {
	return &FileHandler{db: db}
}

const filesBasePath = "files"

func (h *FileHandler) HandleUpload(c *gin.Context) {
	err := c.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	path := filesBasePath + "/" + c.Param("path")
	file, _, err := c.Request.FormFile("file")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	newFile, err := os.Create(path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = io.Copy(newFile, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}

func getAllFilePaths(rootDir string, username string) ([]string, error) {
	var filePaths []string

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			relPath := username + "/" + strings.TrimPrefix(path, rootDir+"/")
			filePaths = append(filePaths, relPath)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return filePaths, nil
}

// 返回用户所有文件组成的列表。客户端接收后根据表的内容逐个请求文件。
func (h *FileHandler) HandleList(c *gin.Context) {
	user := c.Param("user")

	filePaths, err := getAllFilePaths(filesBasePath+"/"+user, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"files": filePaths})
}

func (h *FileHandler) HandleDownload(c *gin.Context) {
	path := filesBasePath + "/" + c.Param("path")

	file, err := os.Open(path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	c.File(path)
}
