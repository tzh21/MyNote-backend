package handler

import (
	"MyNote-backend/utils"
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
// const storageBasePath = "storage"
// var notesBasePath = fmt.Sprintf("%s/notes", storageBasePath)
// var blocksBasePath = fmt.Sprintf("%s/blocks", notesBasePath)

func (h *FileHandler) HandleUploadBlocks(c *gin.Context) {
	username := c.Param("username")
	filename := c.Param("filename")
	savePath := utils.BlocksPath(username, filename)
	HandleUpload(c, savePath)
}

func (h *FileHandler) HandleGetBlocks(c *gin.Context) {
	username := c.Param("username")
	filename := c.Param("filename")
	path := utils.BlocksPath(username, filename)

	c.File(path)
}

func HandleUpload(c *gin.Context, savePath string) {
	// Ensure the directory exists
	dir := filepath.Dir(savePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	// Create the file to save the uploaded content
	newFile, err := os.Create(savePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer newFile.Close()

	// Copy the request body to the new file
	_, err = io.Copy(newFile, c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}

func (h *FileHandler) HandleUploadImage(c *gin.Context) {
	username := c.Param("username")
	filename := c.Param("filename")
	savePath := utils.ImagesPath(username, filename)
	HandleUpload(c, savePath)
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
	HandleUpload(c, savePath)
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

// func HandleUpload(c *gin.Context, savePath string) {
// 	err := c.Request.ParseMultipartForm(10 << 20)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	file, _, err := c.Request.FormFile("file")
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	defer file.Close()

// 	dir := filepath.Dir(savePath)
// 	if _, err := os.Stat(dir); os.IsNotExist(err) {
// 		err := os.MkdirAll(dir, os.ModePerm)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}
// 	}

// 	newFile, err := os.Create(savePath)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	_, err = io.Copy(newFile, file)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
// }

// 获取 dir 下所有笔记文件的相对路径，并在前面添加 username
// func getAllNoteRelPaths(username string) ([]string, error) {
// 	var dir = utils.NotesBasePath(username)

// 	var filePaths []string
// 	// 用户文件夹下默认为笔记分类文件夹
// 	categorys, err := os.ReadDir(dir)
// 	if err != nil {
// 		return nil, err
// 	}

// 	for _, category := range categorys {
// 		if category.IsDir() {
// 			noteFiles, err := os.ReadDir(dir + "/" + category.Name())
// 			if err != nil {
// 				return nil, err
// 			}

// 			for _, file := range noteFiles {
// 				if !file.IsDir() {
// 					relPath := username + "/" + category.Name() + "/" + file.Name()
// 					filePaths = append(filePaths, relPath)
// 				}
// 			}
// 		}
// 	}

// 	return filePaths, nil
// }

// 返回用户所有笔记文件组成的列表。客户端接收后根据表的内容逐个请求笔记文件。
// func (h *FileHandler) HandleList(c *gin.Context) {
// 	user := c.Param("user")

// 	filePaths, err := getAllNoteRelPaths(user)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"files": filePaths})
// }

// func (h *FileHandler) HandleDownload(c *gin.Context) {
// 	path := notesBasePath + "/" + c.Param("path")

// 	file, err := os.Open(path)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	defer file.Close()

// 	c.File(path)
// }
