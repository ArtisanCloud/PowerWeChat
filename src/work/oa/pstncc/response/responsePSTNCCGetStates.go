package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponsePSTNCCGetStates struct {
	*response.ResponseWork

	IsTalked int `json:"istalked"`
	CallTime int `json:"calltime"`
	TalkTime int `json:"talktime"`
	Reason   int `json:"reason"`
}
