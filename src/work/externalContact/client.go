package externalContact

import (
	"fmt"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-socialite/src/response/weCom"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
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
func (comp *Client) GetFollowUsers() (*response.ResponseGetFollowUserList, error) {

	result := &response.ResponseGetFollowUserList{}

	_, err := comp.HttpGet("cgi-bin/externalcontact/get_follow_user_list", nil, nil, result)

	return result, err
}

// 获取外部联系人列表.
// https://open.work.weixin.qq.com/api/doc/90000/90135/92571
func (comp *Client) List(userID string) (*response.ResponseGetList, error) {

	result := &response.ResponseGetList{}

	_, err := comp.HttpGet("cgi-bin/externalcontact/list", &object.StringMap{
		"userid": userID,
	}, nil, result)

	return result, err
}

// 批量获取客户详情.
// https://open.work.weixin.qq.com/api/doc/90000/90135/92994
func (comp *Client) BatchGet(userID string, cursor string, limit int) (*response.ResponseBatchGetByUser, error) {

	result := &response.ResponseBatchGetByUser{}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/batch/get_by_user", &object.StringMap{
		"userid": userID,
		"cursor": cursor,
		"limit":  fmt.Sprintf("%d", limit),
	}, nil, nil, result)

	return result, err
}

// 获取外部联系人详情.
// https://work.weixin.qq.com/api/doc#90000/90135/91556
func (comp *Client) Get(externalUserId string) (*weCom.ResponseGetExternalContact, error) {

	result := &weCom.ResponseGetExternalContact{}

	_, err := comp.HttpGet("cgi-bin/externalcontact/get", &object.StringMap{
		"external_userid": externalUserId,
	}, nil, result)

	return result, err
}

// 修改客户备注信息.
// https://work.weixin.qq.com/api/doc/90000/90135/92115
func (comp *Client) Remark(data *power.HashMap) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/remark", data, nil, nil, result)

	return result, err
}

// 获取离职成员的客户列表.
// https://work.weixin.qq.com/api/doc#90000/90135/91563
func (comp *Client) GetUnassigned(pageID int, pageSize int) (*response.ResponseGetUnassignedList, error) {

	result := &response.ResponseGetUnassignedList{}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/get_unassigned_list", &object.HashMap{
		"page_id":   strconv.Itoa(pageID),
		"page_size": strconv.Itoa(pageSize),
	}, nil, nil, result)

	return result, err
}

// 离职成员的外部联系人再分配.
// https://work.weixin.qq.com/api/doc#90000/90135/91564
func (comp *Client) Transfer(externalUserID []string, handoverUserID string, takeoverUserID string) (*response.ResponseGetTransferedCustomerList, error) {

	result := &response.ResponseGetTransferedCustomerList{}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/transfer_customer", &object.HashMap{
		"handover_userid": handoverUserID,
		"takeover_userid": takeoverUserID,
		"external_userid": externalUserID,
	}, nil, nil, result)

	return result, err
}

// 离职成员的群再分配.
// https://work.weixin.qq.com/api/doc/90000/90135/92127
func (comp *Client) TransferGroupChat(chatIDs []string, newOwner string) (*response.ResponseGroupChatTransfer, error) {

	result := &response.ResponseGroupChatTransfer{}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/groupchat/transfer", &object.HashMap{
		"chat_id_list": chatIDs,
		"new_owner":    newOwner,
	}, nil, nil, result)

	return result, err
}

// 查询客户接替结果.
// https://work.weixin.qq.com/api/doc/90000/90135/94082
func (comp *Client) GetResignedTransferResult(handoverUserID string, takeoverUserID string, cursor string) (*response.ResponseGetTransferedCustomerList, error) {

	result := &response.ResponseGetTransferedCustomerList{}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/resigned/transfer_result?", &object.StringMap{
		"handover_userid": handoverUserID,
		"takeover_userid": takeoverUserID,
		"cursor":          cursor,
	}, nil, nil, result)

	return result, err
}
