package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseWebDriveSpaceCreate struct {
	*response.ResponseWork

	SpaceID string `json:"spaceid"`
}
