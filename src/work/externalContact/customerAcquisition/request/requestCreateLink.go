package request

type Range struct {
	UserList       []string `json:"user_list"`
	DepartmentList []string `json:"department_list"`
}

type PriorityOption struct {
	PriorityType       int      `json:"priority_type"`
	PriorityUseridList []string `json:"priority_userid_list"`
}

type RequestCreateLink struct {
	LinkName       string          `json:"link_name"`   // 链接名称
	Range          *Range          `json:"range"`       // 1,
	SkipVerify     bool            `json:"skip_verify"` // true,
	PriorityOption *PriorityOption `json:"priority_option"`
}

type RequestUpdateLink struct {
	LinkId         string          `json:"link_id"`     // 链接ID
	LinkName       string          `json:"link_name"`   // 链接名称
	Range          *Range          `json:"range"`       // 1,
	SkipVerify     bool            `json:"skip_verify"` // true,
	PriorityOption *PriorityOption `json:"priority_option"`
}

type RequestDeleteLink struct {
	LinkId string `json:"link_id"`
}
