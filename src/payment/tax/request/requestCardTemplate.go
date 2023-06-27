package request

type CustomCell struct {
	Words               string `json:"words"`
	Description         string `json:"description"`
	JumpUrl             string `json:"jump_url"`
	MiniProgramUserName string `json:"miniprogram_user_name"`
	MiniProgramPath     string `json:"miniprogram_path"`
}

type CardTemplateInformation struct {
	PayeeName  string      `json:"payee_name"`
	LogoUrl    string      `json:"logo_url"`
	CustomCell *CustomCell `json:"custom_cell"`
}

type RequestApplyForCardTemplate struct {
	CardAppid               string                   `json:"card_appid"`
	CardTemplateInformation *CardTemplateInformation `json:"card_template_information"`
}
