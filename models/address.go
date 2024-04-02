package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	AddressCompany string
	AddressSpecific string
	CompanyId uint
	Company Company `gorm:"foreignKey:CompanyId"`
	StaffId uint
	Staff Staff `gorm:"foreignKey:StaffId"`
	// staff_id uint
	// company_id uint
}