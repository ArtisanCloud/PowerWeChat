package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseBroadcastGetAssistantList struct {
	*response.ResponseMiniProgram

	List     int `json:"list"`
	Count    int `json:"count"`
	MaxCount int `json:"maxCount"`
}
