package model

import "time"

type Transaction struct {
	BillerID        int        `gorm:"primarykey;autoincrement" json:biller_id`
	AccountID       string     `gorm:"foreignkey" json:"account_id"`
	Amount          float64    `gorm:"column:amount" json:amount`
	TransactionDate *time.Time `gorm:"column:transaction_date" json:"transaction_date"`
}

func (a *Transaction) TableName() string {
	return "transaction"
}
