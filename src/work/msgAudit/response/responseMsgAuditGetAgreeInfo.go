package response

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseMsgAuditGetAgreeInfo struct {
	*response.ResponseWork
	AgreeInfo []*object.HashMap `json:"agreeinfo"`
}
