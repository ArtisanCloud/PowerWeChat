package request

type UserList struct {
	Openid string `json:"openid"`
	Lang   string `json:"lang"`
}

type RequestBatchGetUserInfo struct {
	UserList []*UserList `json:"user_list"`
}
