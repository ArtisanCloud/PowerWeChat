package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type KF struct {
	KFAccount        string `json:"kf_account"`
	KFHeadImgURL     string `json:"kf_headimgurl"`
	KFID             int    `json:"kf_id"`
	KFNick           string `json:"kf_nick"`
	KFWx             string `json:"kf_wx,omitempty"`
	InviteWx         string `json:"invite_wx,omitempty"`
	InviteExpireTime int    `json:"invite_expire_time,omitempty"`
	InviteStatus     string `json:"invite_status,omitempty"`
}

type ResponseList struct {
	response.ResponseOfficialAccount

	KFList []*KF `json:"kf_list"`
}

// ------------------------------------------------------------

type KFOnline struct {
	KFAccount    string `json:"kf_account"`
	Status       int    `json:"status"`
	KFID         string `json:"kf_id"`
	AcceptedCase int    `json:"accepted_case"`
}

type ResponseKFOnlineList struct {
	response.ResponseOfficialAccount

	KfOnlineList []*KFOnline `json:"kf_online_list"`
}
