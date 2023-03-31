package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type Session struct {
	CreateTime int    `json:"createtime"`
	OpenID     string `json:"openid"`
}

type ResponseKFSessionList struct {
	response.ResponseOfficialAccount

	Sessionlist []*Session `json:"sessionlist"`
}

// ---------------------------------------------------

type WaitCase struct {
	LatestTime int    `json:"latest_time"`
	Openid     string `json:"openid"`
}

type ResponseKFSessionWaitCaseList struct {
	response.ResponseOfficialAccount

	Count        int         `json:"count"`
	WaitCaseList []*WaitCase `json:"waitcaselist"`
}

// ---------------------------------------------------

type ResponseKFSessionGet struct {
	response.ResponseOfficialAccount

	CreateTime int    `json:"createtime"`
	KfAccount  string `json:"kf_account"`
}
