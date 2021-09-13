package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseMsgAuditGetAgreeInfo struct {
	*response.ResponseWork
	AgreeInfo []*power.HashMap `json:"agreeinfo"`
}
