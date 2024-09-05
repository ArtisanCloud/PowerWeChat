package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseListLink struct {
	response.ResponseWork

	LinkIdList []string `json:"link_id_list"`
	NextCursor string   `json:"next_cursor"`
}
