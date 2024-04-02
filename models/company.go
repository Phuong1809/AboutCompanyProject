package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name string
	// AddressId uint
	// Address Address `gorm:"foreignKey:AddressId"`
	Staff []Staff `gorm:"many2many:staff_company;"`
}