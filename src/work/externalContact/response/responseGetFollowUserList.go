package response

import "github.com/ArtisanCloud/power-wechat/src/kernel/response"

type ResponseGetFollowUserList struct {
	*response.ResponseWork

	FollowUser []string `json:"follow_user"` // ["zhangsan","tagid2"]
}
