package response

import (
	"github.com/ArtisanCloud/go-socialite/src/models"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseGetUserList struct {
	response.ResponseWork
	//UserList []*response2.RequestUserDetail `json:"userlist"`
	UserList []*models.Employee `json:"userlist"`
}
