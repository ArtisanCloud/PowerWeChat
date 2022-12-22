package comment

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/comment/response"
)

type Client struct {
	*kernel.BaseClient
}

// 打开已群发文章评论
// https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html
func (comp *Client) Open(msgID string, index int) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"msg_data_id": msgID,
		"index":       index,
	}

	_, err := comp.HttpPostJson("cgi-bin/comment/open", params, nil, nil, result)

	return result, err
}

// 关闭已群发文章评论
// https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html
func (comp *Client) Close(msgID string, index int) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"msg_data_id": msgID,
		"index":       index,
	}

	_, err := comp.HttpPostJson("cgi-bin/comment/close", params, nil, nil, result)

	return result, err
}

// 查看指定文章的评论数据
// https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html
func (comp *Client) List(msgID string, index int, begin int, count int, Type int) (*response.ResponseCommentList, error) {

	result := &response.ResponseCommentList{}

	params := &object.HashMap{
		"msg_data_id": msgID,
		"index":       index,
		"begin":       begin,
		"count":       count,
		"type":        Type,
	}

	_, err := comp.HttpPostJson("cgi-bin/comment/list", params, nil, nil, result)

	return result, err
}

// 将评论标记精选（新增接口）
// https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html
func (comp *Client) MarkElect(msgID string, index int, commentID int) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"msg_data_id":     msgID,
		"index":           index,
		"user_comment_id": commentID,
	}

	_, err := comp.HttpPostJson("cgi-bin/comment/markelect", params, nil, nil, result)

	return result, err
}

// 将评论取消精选（新增接口）
// https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html
func (comp *Client) UnmarkElect(msgID string, index int, commentID int) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"msg_data_id":     msgID,
		"index":           index,
		"user_comment_id": commentID,
	}

	_, err := comp.HttpPostJson("cgi-bin/comment/unmarkelect", params, nil, nil, result)

	return result, err
}

// 删除评论（新增接口）
// https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html
func (comp *Client) Delete(msgID string, index int, commentID int) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"msg_data_id":     msgID,
		"index":           index,
		"user_comment_id": commentID,
	}

	_, err := comp.HttpPostJson("cgi-bin/comment/delete", params, nil, nil, result)

	return result, err
}

// 回复评论（新增接口）
// https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html
func (comp *Client) Reply(msgID string, index int, commentID int, content string) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"msg_data_id":     msgID,
		"index":           index,
		"user_comment_id": commentID,
		"content":         content,
	}

	_, err := comp.HttpPostJson("cgi-bin/comment/reply/add", params, nil, nil, result)

	return result, err
}

// 删除回复（新增接口）
// https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html
func (comp *Client) DeleteReply(msgID string, index int, commentID int) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"msg_data_id":     msgID,
		"index":           index,
		"user_comment_id": commentID,
	}

	_, err := comp.HttpPostJson("cgi-bin/comment/reply/delete", params, nil, nil, result)

	return result, err
}
