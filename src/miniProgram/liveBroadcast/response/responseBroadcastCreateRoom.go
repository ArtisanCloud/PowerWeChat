package response

import "github.com/ArtisanCloud/power-wechat/src/kernel/response"

type ResponseBroadcastCreateRoom struct {
	*response.ResponseMiniProgram

	RoomID int `json:"roomId"`
	QRCodeURL string `json:"qrcode_url"`


}
