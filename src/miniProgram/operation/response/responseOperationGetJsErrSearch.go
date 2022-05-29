package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseOperationGetJsErrSearch struct {
	*response.ResponseMiniProgram
	results *power.HashMap `json:"results"`
	Total   int64          `json:"total"`
}
