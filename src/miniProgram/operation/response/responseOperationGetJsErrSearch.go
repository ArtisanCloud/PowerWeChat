package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseOperationGetJsErrSearch struct {
	*response.ResponseMiniProgram
	results *power.HashMap `json:"results"`
	Total   int64          `json:"total"`
}
