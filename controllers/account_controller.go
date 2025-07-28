package controllers

import (
	"internal-transfers/database"
	"internal-transfers/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {
	var acc models.Account
	if err := c.ShouldBindJSON(&acc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	_, err := database.DB.Exec("INSERT INTO accounts (account_id, balance) VALUES ($1, $2)", acc.AccountID, acc.Balance)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Account creation failed"})
		return
	}

	c.Status(http.StatusOK)
}

func GetAccount(c *gin.Context) {
	idStr := c.Param("account_id")
	accountID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
		return
	}

	var acc models.AccountResponse
	err = database.DB.QueryRow("SELECT account_id, balance FROM accounts WHERE account_id = $1", accountID).
		Scan(&acc.AccountID, &acc.Balance)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}

	c.JSON(http.StatusOK, acc)
}
