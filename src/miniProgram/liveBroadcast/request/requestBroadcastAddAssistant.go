package request

import "github.com/ArtisanCloud/power-wechat/src/kernel/power"

type RequestBroadcastAddAssistant struct {
	RoomID int `json:"roomId"`
	Users []power.StringMap `json:"users"`
}
