package response

import  "github.com/ArtisanCloud/power-wechat/src/kernel/response"

type ResponseAddGroupWelcomeTemplate struct {
	*response.ResponseWork

	TemplateID string `json:"template_id"`
}