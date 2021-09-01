package response

import (
	"github.com/ArtisanCloud/go-libs/object"
	response2 "github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseExpressAccountGetAll struct {
	*response2.ResponseMiniProgram

	Count int               `json:"count"`
	List  []*object.HashMap `json:"list"`
}
