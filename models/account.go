package models


type Account struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Username string
	Password string
	Cookie string
}

func NewAccount(Username string, Password string)*Account{
	return &Account{
		Username: Username,
		Password: Password,
	}
}