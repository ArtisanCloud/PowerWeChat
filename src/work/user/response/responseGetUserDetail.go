package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
	response2 "github.com/ArtisanCloud/power-wechat/src/work/user/request"
)

type ResponseGetUserDetail struct {
	*response.ResponseWork

	*response2.RequestUserDetail
}
