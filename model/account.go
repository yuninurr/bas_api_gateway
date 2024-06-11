package model

type Account struct {
	AccountID string ` gorm:"primaryKey" `
	Username  string ` gorm:"column:username" `
	Password  string
	Name      string
}

func (a *Account) TableName() string {
	return "account"
}