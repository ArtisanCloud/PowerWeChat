package response


type KF struct {
	KFAccount        string `json:"kf_account"`
	KFHeadImgURL     string `json:"kf_headimgurl"`
	KFId             string `json:"kf_id"`
	KFNick           string `json:"kf_nick"`
	KFWx             string `json:"kf_wx,omitempty"`
	InviteWx         string `json:"invite_wx,omitempty"`
	InviteExpireTime int    `json:"invite_expire_time,omitempty"`
	InviteStatus     string `json:"invite_status,omitempty"`
}

type ResponseList struct {
	KFList  []*KF `json:"kf_list"`
}

// ------------------------------------------------------------

type KFOnline struct {
	KFAccount    string `json:"kf_account"`
	Status       int    `json:"status"`
	KFID         string `json:"kf_id"`
	AcceptedCase int    `json:"accepted_case"`
}

type ResponseKFOnlineList struct {
	KfOnlineList []*KFOnline `json:"kf_online_list"`
}