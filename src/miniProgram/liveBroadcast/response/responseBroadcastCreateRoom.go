package response

import "github.com/ArtisanCloud/PowerWeChat/src/kernel/response"

type ResponseBroadcastCreateRoom struct {
	*response.ResponseMiniProgram

	RoomID int `json:"roomId"`
	QRCodeURL string `json:"qrcode_url"`


}
