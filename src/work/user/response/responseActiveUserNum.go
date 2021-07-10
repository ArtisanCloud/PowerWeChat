package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseUserActiveCount struct {
	response.ResponseWX
	ActiveCount string `json:"active_cnt"`
}
