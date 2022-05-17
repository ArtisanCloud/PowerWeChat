package response

import "github.com/ArtisanCloud/PowerWeChat/src/kernel/response"

type ResponseGetUserInfo struct {
	Subscribe      int    `json:"subscribe"`
	OpenID         string `json:"openid"`
	Language       string `json:"language"`
	SubscribeTime  int    `json:"subscribe_time"`
	UnionID        string `json:"unionid"`
	Remark         string `json:"remark"`
	GroupID        int    `json:"groupid"`
	TagIDList      []int  `json:"tagid_list"`
	SubscribeScene string `json:"subscribe_scene"`
	QrScene        int    `json:"qr_scene"`
	QrSceneStr     string `json:"qr_scene_str"`
}

type ResponseBatchGetUserInfo struct {
	ResponseGetUserInfo []*ResponseGetUserInfo `json:"user_info_list"`
}

type ResponseGetUserList struct {
	Total int `json:"total"`
	Count int `json:"count"`
	Data  struct {
		OpenID []string `json:"openid"`
	} `json:"data"`
	NextOpenID string `json:"next_openid"`
}

type ResponseChangeOpenID struct {
	*response.ResponseOfficialAccount

	ResultList []struct {
		OriOpenid string `json:"ori_openid"`
		NewOpenid string `json:"new_openid,omitempty"`
		ErrMsg    string `json:"err_msg"`
	} `json:"result_list"`
}