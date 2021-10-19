package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
	response2 "github.com/ArtisanCloud/powerwechat/src/work/user/request"
)

type ResponseGetUserDetail struct {
	*response.ResponseWork

	*response2.RequestUserDetail
}
