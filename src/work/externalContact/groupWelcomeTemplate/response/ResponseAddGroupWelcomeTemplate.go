package response

import  "github.com/ArtisanCloud/powerwechat/src/kernel/response"

type ResponseAddGroupWelcomeTemplate struct {
	*response.ResponseWork

	TemplateID string `json:"template_id"`
}