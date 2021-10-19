package response

import (
	"github.com/ArtisanCloud/PowerSocialite/src/models"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseBatchGetByUser struct {
	*response.ResponseWork

	ExternalContactList []*ResponseExternalContact `json:"external_contact_list"`
	NextCursor          string                     `json:"next_cursor"`
}

type ResponseExternalContact struct {
	ExternalContact *models.ExternalContact `json:"external_contact"`
	FollowInfo      *models.FollowUser      `json:"follow_info"`
}
