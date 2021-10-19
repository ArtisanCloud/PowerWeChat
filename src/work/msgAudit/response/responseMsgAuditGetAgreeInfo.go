package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseMsgAuditGetAgreeInfo struct {
	*response.ResponseWork
	AgreeInfo []*power.HashMap `json:"agreeinfo"`
}
