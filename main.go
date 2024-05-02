package main

import (
	"MyNote-backend/db"
	"MyNote-backend/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	db := db.InitDB()
	r := gin.Default()
	handler.SetupRoutes(r, db)
	r.Run()
}
