package request

type RequestRegisterMiniProgram struct {
	Name               string `json:"name"`
	Code               string `json:"code"`
	CodeType           int    `json:"code_type"`
	LegalPersonaWechat string `json:"legal_persona_wechat"`
	LegalPersonaName   string `json:"legal_persona_name"`
	ComponentPhone     string `json:"component_phone"`
}

type GetRegistrationStatus struct {
	Name               string `json:"name"`
	LegalPersonaWechat string `json:"legal_persona_wechat"`
	LegalPersonaName   string `json:"legal_persona_name"`
}
