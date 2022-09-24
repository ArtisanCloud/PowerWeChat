package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type Title struct {
	Text string `json:"text"`
	Lang string `json:"lang"`
}

type Placeholder struct {
	Text string `json:"text"`
	Lang string `json:"lang"`
}

type Property struct {
	Control     string         `json:"control"`
	ID          string         `json:"id"`
	Title       []*Title       `json:"title"`
	Placeholder []*Placeholder `json:"placeholder"`
	Require     int            `json:"require"`
	UnPrint     int            `json:"un_print"`
}

type Value struct {
	Text string `json:"text"`
	Lang string `json:"lang"`
}

type Option struct {
	Key   string   `json:"key"`
	Value []*Value `json:"value"`
}

type Control struct {
	Property *Property      `json:"property"`
	Config   *power.HashMap `json:"config"`
}

type TemplateContent struct {
	Controls []Control `json:"controls"`
}

type TemplateName struct {
	Text string `json:"text"`
	Lang string `json:"lang"`
}

type ResponseApprovalTemplate struct {
	*response.ResponseWork

	TemplateNames   []*TemplateName  `json:"template_names"`
	TemplateContent *TemplateContent `json:"template_content"`
}
