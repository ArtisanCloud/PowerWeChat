package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseAuthGetPaidUnionID struct {
	UnionID string `json:"unionid"`

	*response.ResponseMiniProgram
}
