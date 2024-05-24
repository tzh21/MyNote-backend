package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

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

const storageBasePath = "storage"

var notesBasePath = fmt.Sprintf("%s/note", storageBasePath)

func BlocksBasePath(username string) string {
	return fmt.Sprintf("%s/%s/blocks", notesBasePath, username)
}

func BlocksPath(username, filename string) string {
	return fmt.Sprintf("%s/%s", BlocksBasePath(username), filename)
}

func ImagesBasePath(username string) string {
	return fmt.Sprintf("%s/%s/images", notesBasePath, username)
}

func ImagesPath(username, filename string) string {
	return fmt.Sprintf("%s/%s", ImagesBasePath(username), filename)
}

func AudioBasePath(username string) string {
	return fmt.Sprintf("%s/%s/audio", notesBasePath, username)
}

func AudioPath(username, filename string) string {
	return fmt.Sprintf("%s/%s", AudioBasePath(username), filename)
}

var profileBasePath = fmt.Sprintf("%s/profile", storageBasePath)

func AvatarBasePath(username string) string {
	return fmt.Sprintf("%s/%s/avatar", profileBasePath, username)
}

func AvatarPath(username, filename string) string {
	return fmt.Sprintf("%s/%s", AvatarBasePath(username), filename)
}
