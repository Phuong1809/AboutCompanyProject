package models

import "gorm.io/gorm"

type QR struct {
	gorm.Model
	FilePatch string
	StaffId uint
	Staff Staff `gorm:"foreignKey:StaffId"`
}