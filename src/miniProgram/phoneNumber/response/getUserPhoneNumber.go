package response

import "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"

type ResponseGetUserPhoneNumber struct {
	*response.ResponseMiniProgram

	PhoneInfo *struct {
		PhoneNumber     string `json:"phoneNumber,omitempty"`
		PurePhoneNumber string `json:"purePhoneNumber,omitempty"`
		CountryCode     string `json:"countryCode,omitempty"`
		Watermark       *struct {
			Timestamp int    `json:"timestamp"`
			AppID     string `json:"appid,omitempty"`
		} `json:"watermark,omitempty"`
	} `json:"phone_info,omitempty"`
}
