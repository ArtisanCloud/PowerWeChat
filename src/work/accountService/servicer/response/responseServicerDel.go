package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseServicerDel struct {
	response.ResponseWork

	ResultList []*power.HashMap `json:"result_list"`
}
