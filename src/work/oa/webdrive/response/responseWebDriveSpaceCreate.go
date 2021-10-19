package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseWebDriveSpaceCreate struct {
	*response.ResponseWork

	SpaceID string `json:"spaceid"`
}
