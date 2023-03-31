package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseDataCubeVisitInfo struct {
	response.ResponseMiniProgram

	RefDate    string           `json:"ref_date"`
	VisitUVNew []*power.HashMap `json:"visit_uv_new"`
	VisitUV    []*power.HashMap `json:"visit_uv"`
}
