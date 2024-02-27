package main

import (
	"github.com/gin-gonic/gin"
	"salawat/api/api/v1/routes"
	"salawat/api/database"
)

func main() {
	database.Connect()
	r := gin.Default()
	routes.Setup(r)
	err := r.Run("localhost:8080")
	if err != nil {
		return
	}
}
