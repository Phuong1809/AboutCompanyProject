package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name string `json:"name"`
	Address []Address `json:"address"`
	Staff []Staff `json:"staff" gorm:"many2many:staff_company;"`
}

func NewCompany(Name string) *Company{
	return &Company{
		Name: Name,
	}
}