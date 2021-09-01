package response

import "github.com/ArtisanCloud/power-wechat/src/kernel/response"

type ResponseExpressGetPrinter struct {

	*response.ResponseMiniProgram

	Count     string   `json:"count"`
	Openid    []string `json:"openid"`
	TagIDList []string `json:"tagid_list"`
}


