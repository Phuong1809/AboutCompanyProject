package form



type StaffForm struct {
	StaffName   string `json:"staff_name"`
	// FilePathAvatar string `json:"file_path_avatar"`
	Role         string `json:"role"`
	NameCompany string	`json:"name_company"`
	TelePhone    string `json:"tele_phone"`
	MobilePhone  string	`json:"mobile_phone"`
	Fax          string	`json:"fax"`
	Email        string	`json:"email"`
	AddressID uint `json:"address_id"`
}


