package response

import (
	"github.com/ArtisanCloud/PowerSocialite/v2/src/models"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseGetUserList struct {
	*response.ResponseWork

	//UserList []*response2.RequestUserDetail `json:"userlist"`
	UserList []*models.Employee `json:"userlist"`
}
