package request

type RequestBatchGetUserInfo struct {
	UserList []struct {
		Openid string `json:"openid"`
		Lang   string `json:"lang"`
	} `json:"user_list"`
}
