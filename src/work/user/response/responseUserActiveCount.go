package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseJoinCode struct {
	*response.ResponseWork

	JoinCode string `json:"join_qrcode"`
}
