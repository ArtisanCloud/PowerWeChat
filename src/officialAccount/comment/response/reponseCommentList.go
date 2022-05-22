package response

import "github.com/ArtisanCloud/PowerWeChat/src/kernel/response"

type Reply struct {
	Content    string `json:"content"`
	CreateTime int    `json:"create_time"`
}

type Comment struct {
	UserCommentID int    `json:"user_comment_id"`
	OpenID        string `json:"openid"`
	CreateTime    int    `json:"create_time"`
	Content       string `json:"content"`
	CommentType   int    `json:"comment_type"`
	Reply         *Reply `json:"reply"`
}

type ResponseCommentList struct {
	*response.ResponseOfficialAccount

	Total   int        `json:"total"`
	Comment []*Comment `json:"comment"`
}
