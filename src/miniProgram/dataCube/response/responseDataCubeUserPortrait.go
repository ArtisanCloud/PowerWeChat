package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseDataCubeUserPortrait struct {
	*response.ResponseMiniProgram

	RefDate    string         `json:"ref_date"`
	VisitUVNew *power.HashMap `json:"visit_uv_new"`
	VisitUV    *power.HashMap `json:"visit_uv"`
}
