package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseMomentGetMomentComments struct {
	*response.ResponseWork

	CommentList []*power.HashMap `json:"comment_list"`
	LikeList    []*power.HashMap `json:"like_list"`
	NextCursor  string           `json:"next_cursor"`
}
