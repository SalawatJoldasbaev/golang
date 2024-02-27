package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"salawat/api/api/v1/routes"
	"salawat/api/common/middleware"
	"salawat/api/config"
)

func main() {
	var (
		db = config.DBSetup()
	)
	defer config.DBClose(db)
	server := gin.Default()
	server.Use(
		middleware.CORSMiddleware(),
	)
	routes.Setup(server)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := server.Run(":" + port)
	if err != nil {
		fmt.Println("Server failed to start: ", err)
		return
	}
}
