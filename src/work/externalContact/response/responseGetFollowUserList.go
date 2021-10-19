package response

import "github.com/ArtisanCloud/powerwechat/src/kernel/response"

type ResponseGetFollowUserList struct {
	*response.ResponseWork

	FollowUser []string `json:"follow_user"` // ["zhangsan","tagid2"]
}
