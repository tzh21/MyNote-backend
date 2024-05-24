package utils

import "fmt"

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
