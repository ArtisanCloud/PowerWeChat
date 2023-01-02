package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseLinkCorpGetDepartmentList struct {
	response.ResponseWork

	department_list []*power.HashMap `json:"department_list"`
}
