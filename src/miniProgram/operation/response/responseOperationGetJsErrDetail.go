package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseOperationGetJsErrDetail struct {
	*response.ResponseMiniProgram
	Success bool           `json:"success"`
	OpenID  string         `json:"openid"`
	Data   []*power.HashMap `json:"data"`
}
