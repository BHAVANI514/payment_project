package main

import (
	"internal-transfers/database"
	"internal-transfers/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8004")
}
