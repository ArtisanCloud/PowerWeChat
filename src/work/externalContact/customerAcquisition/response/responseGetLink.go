package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type Range struct {
	UserList       []string `json:"user_list"`
	DepartmentList []string `json:"department_list"`
}

type PriorityOption struct {
	PriorityType       int      `json:"priority_type"`
	PriorityUseridList []string `json:"priority_userid_list"`
}

type ResponseGetLink struct {
	response.ResponseWork

	Link           *Link           `json:"link"`
	Range          *Range          `json:"range"`
	PriorityOption *PriorityOption `json:"priority_option"`
}
