package models

type Account struct {
	AccountID int    `json:"account_id" db:"account_id"`
	Balance   string `json:"initial_balance" db:"balance"`
}

type AccountResponse struct {
	AccountID int    `json:"account_id"`
	Balance   string `json:"balance"`
}
