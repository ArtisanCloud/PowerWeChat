package response

import "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"

type ResponseExpressGetQuota struct {
	*response.ResponseMiniProgram

	QuotaNum string `json:"quota_num"`
}
