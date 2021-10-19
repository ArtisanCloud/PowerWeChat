package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
	response2 "github.com/ArtisanCloud/PowerWeChat/src/work/user/request"
)

type ResponseGetUserDetail struct {
	*response.ResponseWork

	*response2.RequestUserDetail
}
