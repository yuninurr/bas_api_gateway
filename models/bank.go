package models

type Bank struct {
	Bank_code string ` gorm:"primaryKey" `
	Name      string ` gorm:"column:name" `
	Address   string
}

func (a *Bank) TableName() string {
	return "bank"
}