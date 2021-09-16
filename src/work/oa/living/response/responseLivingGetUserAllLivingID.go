package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseLivingGetUserAllLivingID struct {
	*response.ResponseWork

	NextCursor   string   `json:"next_cursor"`
	LivingIDList []string `json:"livingid_list"`
}
