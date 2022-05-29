package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseDataCubeVisit struct {
	*response.ResponseMiniProgram

	RefDate string           `json:"ref_date"`
	List    []*power.HashMap `json:"list"`
}
