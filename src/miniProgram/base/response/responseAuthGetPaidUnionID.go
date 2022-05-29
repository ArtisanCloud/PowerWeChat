package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseAuthGetPaidUnionID struct {
	UnionID string `json:"unionid"`

	*response.ResponseMiniProgram
}
