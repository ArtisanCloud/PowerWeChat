package externalContact

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-wechat/src/kernel"
	"github.com/ArtisanCloud/go-wechat/src/work/externalContact/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

// https://work.weixin.qq.com/api/doc#90000/90135/91554

// 获取配置了客户联系功能的成员列表.
// https://open.work.weixin.qq.com/api/doc/90000/90135/92571
func (comp *Client) GetFollowUsers() *response.ResponseGetFollowUserList {

	result := &response.ResponseGetFollowUserList{}

	comp.HttpGet("cgi-bin/externalcontact/get_follow_user_list", nil, result)

	return result
}

// 获取外部联系人列表.
// https://open.work.weixin.qq.com/api/doc/90000/90135/92571
func (comp *Client) List(userID string) *response.ResponseGetList {

	result := &response.ResponseGetList{}

	comp.HttpGet("cgi-bin/externalcontact/list", &object.StringMap{
		"userid": userID,
	}, result)

	return result
}

// 批量获取客户详情.
// https://open.work.weixin.qq.com/api/doc/90000/90135/92994
func (comp *Client) BatchGet(userID string, cursor string, limit string) *response.ResponseBatchGetByUser {

	result := &response.ResponseBatchGetByUser{}

	comp.HttpGet("cgi-bin/externalcontact/batch/get_by_user", &object.StringMap{
		"userid": userID,
		"cursor": cursor,
		"limit":  limit,
	}, result)

	return result
}

// 获取外部联系人详情.
// https://work.weixin.qq.com/api/doc#90000/90135/91556
func (comp *Client) Get(externalUserId string) *response.ResponseGetExternalContact {

	result := &response.ResponseGetExternalContact{}

	comp.HttpGet("cgi-bin/externalcontact/get", &object.StringMap{
		"external_userid": externalUserId,
	}, result)

	return result
}
