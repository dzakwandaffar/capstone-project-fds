package model

type Account struct {
	AccountID string `gorm:"primarykey"`
	Name      string
	Username  string `gorm:"column-username"`
	Password  string
	Address   string
}

func (a *Account) TableName() string {
	return "account"
}
