package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type Watermark struct {
	Timestamp int    `json:"timestamp"`
	AppID     string `json:"appid,omitempty"`
}

type PhoneInfo struct {
	PhoneNumber     string     `json:"phoneNumber,omitempty"`
	PurePhoneNumber string     `json:"purePhoneNumber,omitempty"`
	CountryCode     string     `json:"countryCode,omitempty"`
	Watermark       *Watermark `json:"watermark,omitempty"`
}

type ResponseGetUserPhoneNumber struct {
	response.ResponseMiniProgram

	ErrCode   int        `json:"errcode"`
	ErrMsg    string     `json:"errmsg"`
	PhoneInfo *PhoneInfo `json:"phone_info,omitempty"`
}
