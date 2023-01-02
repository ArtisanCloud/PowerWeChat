package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseWebDriveSpaceCreate struct {
	response.ResponseWork

	SpaceID string `json:"spaceid"`
}
