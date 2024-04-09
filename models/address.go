package models


type Address struct {
	ID uint `gorm:"primaryKey" json:"id"`
	AddressCompany string `json:"address_company"`
	AddressSpecific string	`json:"address_specific"`
	CompanyID uint `json:"company_id"`
	Company Company `json:"company" gorm:"foreignKey:CompanyID"`
	Staff []Staff `json:"staff"`
}

func NewAddresss(AddressCompany string, AddressSpecific string, CompanyID uint) *Address{
	return &Address{
		AddressCompany: AddressCompany,
		AddressSpecific: AddressSpecific,
		CompanyID: CompanyID,
	}
}