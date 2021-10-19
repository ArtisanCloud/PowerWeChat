package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseUserActiveCount struct {
	*response.ResponseWork

	ActiveCount string `json:"active_cnt"`
}
