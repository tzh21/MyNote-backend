package handler

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FileHandler struct {
	db *gorm.DB
}

func NewFileHandler(db *gorm.DB) *FileHandler {
	return &FileHandler{db: db}
}

// 笔记文件存储的根目录
const notesBasePath = "storage/note"

func (h *FileHandler) HandleUpload(c *gin.Context) {
	err := c.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	path := notesBasePath + "/" + c.Param("path")
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

// 获取 dir 下所有笔记文件的相对路径，并在前面添加 username
func getAllNoteRelPaths(username string) ([]string, error) {
	var dir = notesBasePath + "/" + username

	var filePaths []string
	// 用户文件夹下默认为笔记分类文件夹
	categorys, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, category := range categorys {
		if category.IsDir() {
			noteFiles, err := os.ReadDir(dir + "/" + category.Name())
			if err != nil {
				return nil, err
			}

			for _, file := range noteFiles {
				if !file.IsDir() {
					relPath := username + "/" + category.Name() + "/" + file.Name()
					filePaths = append(filePaths, relPath)
				}
			}
		}
	}

	return filePaths, nil
}

// 返回用户所有笔记文件组成的列表。客户端接收后根据表的内容逐个请求笔记文件。
func (h *FileHandler) HandleList(c *gin.Context) {
	user := c.Param("user")

	filePaths, err := getAllNoteRelPaths(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"files": filePaths})
}

func (h *FileHandler) HandleDownload(c *gin.Context) {
	path := notesBasePath + "/" + c.Param("path")

	file, err := os.Open(path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	c.File(path)
}
