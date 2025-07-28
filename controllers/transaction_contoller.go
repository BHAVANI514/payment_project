package controllers

import (
	"internal-transfers/database"
	"internal-transfers/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SubmitTransaction(c *gin.Context) {
	var tx models.Transaction

	// Bind and validate JSON
	if err := c.ShouldBindJSON(&tx); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction data"})
		return
	}

	// Convert Amount (string → float64)
	amount, err := strconv.ParseFloat(tx.Amount, 64)
	if err != nil || amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount format"})
		return
	}

	// Begin DB transaction
	db := database.DB
	sqlTx, err := db.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Transaction start failed"})
		return
	}
	defer sqlTx.Rollback()

	// Deduct from source account
	if res, err := sqlTx.Exec(
		"UPDATE accounts SET balance = balance - $1 WHERE account_id = $2 AND balance >= $1",
		amount, tx.SourceAccountID,
	); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to deduct amount"})
		return
	} else if rows, _ := res.RowsAffected(); rows == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient funds or invalid source account"})
		return
	}

	// Credit destination account
	if res, err := sqlTx.Exec(
		"UPDATE accounts SET balance = balance + $1 WHERE account_id = $2",
		amount, tx.DestinationAccountID,
	); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to credit destination account"})
		return
	} else if rows, _ := res.RowsAffected(); rows == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid destination account"})
		return
	}

	// Log the transaction
	if _, err := sqlTx.Exec(
		"INSERT INTO transactions (source_account_id, destination_account_id, amount) VALUES ($1, $2, $3)",
		tx.SourceAccountID, tx.DestinationAccountID, amount,
	); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log transaction"})
		return
	}

	// Commit transaction
	if err := sqlTx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction"})
		return
	}

	// Success response
	c.Status(http.StatusNoContent)
}
