package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseInvoiceGetInfo struct {
	response.ResponseWork

	CardID    string         `json:"card_id"`
	BeginTime string         `json:"begin_time"`
	EndTime   string         `json:"end_time"`
	OpenID    string         `json:"openid"`
	Type      string         `json:"type"`
	Payee     string         `json:"payee"`
	Detail    string         `json:"detail"`
	UserInfo  *power.HashMap `json:"user_info"`
}
