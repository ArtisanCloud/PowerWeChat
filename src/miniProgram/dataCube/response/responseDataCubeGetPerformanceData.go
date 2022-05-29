package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseDataCubeGetPerformanceData struct {
	Data *power.HashMap `json:"data"`

	*response.ResponseMiniProgram
}
