package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseGetPaidUnionID struct {
	UnionID string `json:"unionid"`

	*response.ResponseMiniProgram
}
