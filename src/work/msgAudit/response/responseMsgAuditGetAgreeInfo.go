package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseMsgAuditGetAgreeInfo struct {
	*response.ResponseWork
	AgreeInfo []*power.HashMap `json:"agreeinfo"`
}
