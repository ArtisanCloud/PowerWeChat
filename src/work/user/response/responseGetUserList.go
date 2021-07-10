package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
	response2 "github.com/ArtisanCloud/power-wechat/src/work/user/request"
)

type ResponseGetUserList struct {
	response.ResponseWX
	UserList []*response2.RequestUserDetail `json:"userlist"`

}

