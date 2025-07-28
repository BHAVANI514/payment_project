package routes

import (
	"internal-transfers/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/accounts", controllers.CreateAccount)
	router.GET("/accounts/:account_id", controllers.GetAccount)
	router.POST("/transactions", controllers.SubmitTransaction)
}
