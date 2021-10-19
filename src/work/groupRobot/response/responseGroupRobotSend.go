package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseGroupRobotSend struct {
	*response.ResponseWork

	Type      string `json:"type"`
	MediaID   string `json:"media_id"`
	CreatedAt string `json:"created_at"`
}
