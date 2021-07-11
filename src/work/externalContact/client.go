package externalContact

import (
	"fmt"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	response2 "github.com/ArtisanCloud/power-wechat/src/kernel/response"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/response"
	"strconv"
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
func (comp *Client) BatchGet(userID string, cursor string, limit int) *response.ResponseBatchGetByUser {

	result := &response.ResponseBatchGetByUser{}

	comp.HttpPostJson("cgi-bin/externalcontact/batch/get_by_user", &object.StringMap{
		"userid": userID,
		"cursor": cursor,
		"limit":  fmt.Sprintf("%d", limit),
	}, nil, result)

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

// 修改客户备注信息.
// https://work.weixin.qq.com/api/doc/90000/90135/92115
func (comp *Client) Remark(data *object.HashMap) *response2.ResponseWX {

	result := &response2.ResponseWX{}

	comp.HttpPostJson("cgi-bin/externalcontact/remark", data, nil, result)

	return result
}

// 获取离职成员的客户列表.
// https://work.weixin.qq.com/api/doc#90000/90135/91563
func (comp *Client) GetUnassigned(pageID int, pageSize int) *response.ResponseGetUnassignedList {

	result := &response.ResponseGetUnassignedList{}

	comp.HttpPostJson("cgi-bin/externalcontact/get_unassigned_list", object.HashMap{
		"page_id":   strconv.Itoa(pageID),
		"page_size": strconv.Itoa(pageSize),
	}, nil, result)

	return result
}

// 离职成员的外部联系人再分配.
// https://work.weixin.qq.com/api/doc#90000/90135/91564
func (comp *Client) Transfer(externalUserID []string, handoverUserID string, takeoverUserID string) *response.ResponseGetTransferedCustomerList {

	result := &response.ResponseGetTransferedCustomerList{}

	comp.HttpPostJson("cgi-bin/externalcontact/transfer_customer", object.HashMap{
		"handover_userid": handoverUserID,
		"takeover_userid": takeoverUserID,
		"external_userid": externalUserID,
	}, nil, result)

	return result
}

// 离职成员的群再分配.
// https://work.weixin.qq.com/api/doc/90000/90135/92127
func (comp *Client) TransferGroupChat(chatIDs []string, newOwner string) *response.ResponseGroupChatTransfer {

	result := &response.ResponseGroupChatTransfer{}

	comp.HttpPostJson("cgi-bin/externalcontact/groupchat/transfer", object.HashMap{
		"chat_id_list": chatIDs,
		"new_owner":    newOwner,
	}, nil, result)

	return result
}

// 查询客户接替结果.
// https://work.weixin.qq.com/api/doc/90000/90135/94082
func (comp *Client) GetResignedTransferResult(handoverUserID string, takeoverUserID string, cursor string) *response.ResponseGetTransferedCustomerList {

	result := &response.ResponseGetTransferedCustomerList{}

	comp.HttpPostJson("cgi-bin/externalcontact/resigned/transfer_result?", &object.StringMap{
		"handover_userid": handoverUserID,
		"takeover_userid": takeoverUserID,
		"cursor":          cursor,
	}, nil, result)

	return result
}
