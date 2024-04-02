package database

import (
	"test/models"
	"gorm.io/gorm"
)



func ORM(db *gorm.DB) error{
	er:=db.AutoMigrate(&models.Account{}, &models.Address{},&models.Company{}, &models.QR{}, &models.Staff{})
	if er!=nil{
		return er
	}
	return er
}