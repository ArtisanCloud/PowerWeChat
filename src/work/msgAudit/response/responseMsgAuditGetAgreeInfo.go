package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseMsgAuditGetAgreeInfo struct {
	*response.ResponseWork
	AgreeInfo []*power.HashMap `json:"agreeinfo"`
}
