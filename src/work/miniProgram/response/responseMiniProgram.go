package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type ResponseSession struct {
	response.ResponseMiniProgram

	CorpID     string `json:"corpid"`
	UserID     string `json:"userid"`
	DeviceID   string `json:"deviceid"`
	SessionKey string `json:"session_key"`
}
