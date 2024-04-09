package models


type Staff struct {
	ID		   uint `gorm:"primaryKey" json:"id"`
	StaffName   string `json:"staff_name"`
	FilePathAvatar string `json:"file_path_avatar"`
	Role         string `json:"role"`
	NameCompany string	`json:"name_company"`
	TelePhone    string `json:"tele_phone"`
	MobilePhone  string	`json:"mobile_phone"`
	Fax          string	`json:"fax"`
	Email        string	`json:"email"`
	AddressID uint `json:"address_id"`
	Address Address `json:"address" gorm:"foreignKey:AddressID"`
	QR []QR `json:"qr"`
	Company      []Company `json:"company" gorm:"many2many:staff_company;"`

}

func NewStaff(StaffName string, Role string, 
	NameCompany string, TelePhone string, MobilePhone string, Fax string, 
	Email string, AddressID uint) *Staff{
	return &Staff{
		StaffName: StaffName,
		Role: Role,
		NameCompany: NameCompany,
		TelePhone: TelePhone,
		MobilePhone: MobilePhone,
		Fax: Fax,
		Email: Email,
		AddressID: AddressID,
	}
}