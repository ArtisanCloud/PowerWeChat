package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseJoinCode struct {
	*response.ResponseWork

	JoinCode string `json:"join_qrcode"`
}
