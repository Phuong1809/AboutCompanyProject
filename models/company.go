package models


type Company struct {
    ID  uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Address []Address `json:"address"`
	Staff []Staff `json:"staff" gorm:"many2many:staff_company;"`
}

func NewCompany(Name string) *Company{
	return &Company{
		Name: Name,
	}
}