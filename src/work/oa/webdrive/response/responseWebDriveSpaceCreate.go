package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseWebDriveSpaceCreate struct {
	*response.ResponseWork

	SpaceID string `json:"spaceid"`
}
