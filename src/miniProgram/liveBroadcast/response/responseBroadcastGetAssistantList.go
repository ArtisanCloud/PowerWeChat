package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseBroadcastGetAssistantList struct {
	response.ResponseMiniProgram
	List []struct {
		Timestamp int    `json:"timestamp"`
		HeadImg   string `json:"headimg,omitempty"`
		Nickname  string `json:"nickname,omitempty"`
		Alias     string `json:"alias,omitempty"`
		OpenID    string `json:"openid,omitempty"`
	} `json:"list"`
	Count    int `json:"count"`
	MaxCount int `json:"maxCount"`
}
