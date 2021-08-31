package response

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseDataCubeVisitInfo struct {
	*response.ResponseMiniProgram

	RefDate    string            `json:"ref_date"`
	VisitUVNew []*object.HashMap `json:"visit_uv_new"`
	VisitUV    []*object.HashMap `json:"visit_uv"`
}
