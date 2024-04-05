package form


type AdressForm struct{
	AddressCompany string `json:"address_company"`
	AddressSpecific string	`json:"address_specific"`
	CompanyID uint `json:"company_id"`
}