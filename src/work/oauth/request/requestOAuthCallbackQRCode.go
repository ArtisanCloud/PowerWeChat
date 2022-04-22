package request

type ParaOAuthCallbackQRCode struct {
	Code  string `form:"code" json:"code" xml:"code" binding:"required"`
	AppID string `form:"appid" json:"appid" xml:"appid"`
	State string `form:"state" json:"state" xml:"state"`
}

