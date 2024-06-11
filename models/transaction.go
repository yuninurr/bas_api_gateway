package models

import "time"

type Transaction struct {
	Id              int ` gorm:"primaryKey" `
	AccountID       string
	BankID          string ` gorm:"column:bank_id" `
	Amount          int
	TransactionDate *time.Time
}

func (a *Transaction) TableName() string {
	return "transaction"
}