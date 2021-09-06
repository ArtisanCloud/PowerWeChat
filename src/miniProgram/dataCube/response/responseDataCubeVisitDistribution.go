package response

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseDataCubeVisit struct {
	*response.ResponseMiniProgram

	RefDate string            `json:"ref_date"`
	List    []*object.HashMap `json:"list"`
}
