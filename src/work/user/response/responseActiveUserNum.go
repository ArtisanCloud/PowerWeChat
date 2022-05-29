package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseUserActiveCount struct {
	*response.ResponseWork

	ActiveCount string `json:"active_cnt"`
}
