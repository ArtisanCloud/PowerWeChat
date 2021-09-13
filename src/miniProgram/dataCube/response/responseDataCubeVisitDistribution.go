package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseDataCubeVisit struct {
	*response.ResponseMiniProgram

	RefDate string            `json:"ref_date"`
	List    []*power.HashMap `json:"list"`
}
