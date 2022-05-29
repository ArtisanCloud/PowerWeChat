package response

import "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"

type ResponseAddGroupWelcomeTemplate struct {
	*response.ResponseWork

	TemplateID string `json:"template_id"`
}
