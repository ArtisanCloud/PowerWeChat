package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type ResponseRegister struct {
	response.ResponseOpenPlatform

	AppID             string `json:"appid"`
	AuthorizationCode string `json:"authorization_code"`
	IsWxVerifySucc    string `json:"is_wx_verify_succ"`
	IsLinkSucc        string `json:"is_link_succ"`
}
