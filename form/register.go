package form

type FormRegister struct{
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
	RePassWord string `json:"re_pass_word"`
}