package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseServicerDel struct {
	*response.ResponseWork

	ResultList []*power.HashMap `json:"result_list"`
}
