package response

import "github.com/ArtisanCloud/PowerWeChat/src/kernel/response"

type Industry struct {
	FirstClass  string `json:"first_class"`
	SecondClass string `json:"second_class"`
}

type ResponseTemplateIndustry struct {
	*response.ResponseOfficialAccount

	PrimaryIndustry   *Industry `json:"primary_industry"`
	SecondaryIndustry *Industry `json:"secondary_industry"`
}

type ResponseTemplate struct {
	*response.ResponseOfficialAccount

	TemplateId string `json:"template_id"`
}

type Template struct {
	TemplateID      string `json:"template_id"`
	Title           string `json:"title"`
	PrimaryIndustry string `json:"primary_industry"`
	DeputyIndustry  string `json:"deputy_industry"`
	Content         string `json:"content"`
	Example         string `json:"example"`
}

type ResponseTemplateGetPrivate struct {
	*response.ResponseOfficialAccount

	TemplateList []*Template `json:"template_list"`
}

type ResponseTemplateSend struct {
	*response.ResponseOfficialAccount

	MsgID int `json:"msgid"`
}
