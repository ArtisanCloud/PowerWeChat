package request

type RequestBroadcastGetRoleList struct {
	Role    int    `json:"role"`
	Offset  int    `json:"offset"`
	Limit   int    `json:"limit"`
	Keyword string `json:"keyword,omitempty"`
}
