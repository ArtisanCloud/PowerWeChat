package response

import (
	"github.com/ArtisanCloud/PowerSocialite/src/models"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseGetUserList struct {
	*response.ResponseWork

	//UserList []*response2.RequestUserDetail `json:"userlist"`
	UserList []*models.Employee `json:"userlist"`
}
