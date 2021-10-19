package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseLivingGetLivingShareInfo struct {
	*response.ResponseWork

	LivingID              string `json:"livingid"`
	ViewerUserID          string `json:"viewer_userid"`
	ViewerExternalUserID  string `json:"viewer_external_userid"`
	InvitorUserid         string `json:"invitor_userid"`
	InvitorExternalUserID string `json:"invitor_external_userid"`
}
