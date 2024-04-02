package models

import (
	"gorm.io/gorm"
)

type Staff struct {
	gorm.Model
	StaffName   string
	Role         string
	NameCompany string
	// AddressId	uint
	// Address      Address `gorm:"foreignKey:AddressId"`
	TelePhone    string
	MobilePhone  string
	Fax          string
	Email        string
	// QrId        uint
	// QR           QR        `gorm:"foreignKey:QrId"`
	Company      []Company `gorm:"many2many:staff_company;"`
}
