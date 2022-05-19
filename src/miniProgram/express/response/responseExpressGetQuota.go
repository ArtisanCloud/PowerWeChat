package response

import "github.com/ArtisanCloud/PowerWeChat/src/kernel/response"

type ResponseExpressGetQuota struct {
	*response.ResponseMiniProgram

	QuotaNum string `json:"quota_num"`
}
