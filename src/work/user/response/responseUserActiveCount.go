package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseJoinCode struct {
	response.ResponseWX
	JoinCode string `json:"join_qrcode"`
}
