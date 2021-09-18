package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseLinkCorpGetDepartmentList struct {
	*response.ResponseWork

	department_list []*power.HashMap `json:"department_list"`

}
