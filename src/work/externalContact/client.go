package externalContact

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerSocialite/v3/src/response/weCom"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	response3 "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/groupChat/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/response"
	"strconv"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(&app, nil)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient,
	}, nil
}

// 获取配置了客户联系功能的成员列表.
// https://developer.work.weixin.qq.com/document/path/92571
func (comp *Client) GetFollowUsers(ctx context.Context) (*response.ResponseGetFollowUserList, error) {

	result := &response.ResponseGetFollowUserList{}

	_, err := comp.BaseClient.HttpGet(ctx, "cgi-bin/externalcontact/get_follow_user_list", nil, nil, result)

	return result, err
}

// 获取外部联系人列表.
// https://developer.work.weixin.qq.com/document/path/92113
func (comp *Client) List(ctx context.Context, userID string) (*response.ResponseGetList, error) {

	result := &response.ResponseGetList{}

	_, err := comp.BaseClient.HttpGet(ctx, "cgi-bin/externalcontact/list", &object.StringMap{
		"userid": userID,
	}, nil, result)

	return result, err
}

// 获取外部联系人详情.
// https://developer.work.weixin.qq.com/document/path/92114
func (comp *Client) Get(ctx context.Context, externalUserID string, cursor string) (*weCom.ResponseGetExternalContact, error) {

	result := &weCom.ResponseGetExternalContact{}

	_, err := comp.BaseClient.HttpGet(ctx, "cgi-bin/externalcontact/get", &object.StringMap{
		"external_userid": externalUserID,
		"cursor":          cursor,
	}, nil, result)

	return result, err
}

// 批量获取客户详情.
// https://developer.work.weixin.qq.com/document/path/92994
func (comp *Client) BatchGet(ctx context.Context, userID []string, cursor string, limit int) (*response.ResponseBatchGetByUser, error) {

	result := &response.ResponseBatchGetByUser{}

	options := &object.HashMap{
		"userid_list": userID,
		"cursor":      cursor,
		"limit":       limit,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/batch/get_by_user", options, nil, nil, result)

	return result, err
}

// 修改客户备注信息.
// https://developer.work.weixin.qq.com/document/path/92115
func (comp *Client) Remark(ctx context.Context, data *request.RequestExternalContactRemark) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/remark", data, nil, nil, result)

	return result, err
}

// 获取待分配的离职成员列表
// https://developer.work.weixin.qq.com/document/path/92124
func (comp *Client) GetUnassigned(ctx context.Context, pageID int, pageSize int) (*response.ResponseGetUnassignedList, error) {

	result := &response.ResponseGetUnassignedList{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/get_unassigned_list", &object.HashMap{
		"page_id":   strconv.Itoa(pageID),
		"page_size": strconv.Itoa(pageSize),
	}, nil, nil, result)

	return result, err
}

// 分配离职成员的客户
// https://developer.work.weixin.qq.com/document/path/94081
func (comp *Client) Transfer(ctx context.Context, externalUserID []string, handoverUserID string, takeoverUserID string) (*response.ResponseGetTransferedCustomerList, error) {

	result := &response.ResponseGetTransferedCustomerList{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/transfer_customer", &object.HashMap{
		"handover_userid": handoverUserID,
		"takeover_userid": takeoverUserID,
		"external_userid": externalUserID,
	}, nil, nil, result)

	return result, err
}

// 分配离职成员的客户群
// https://developer.work.weixin.qq.com/document/path/92127
func (comp *Client) TransferGroupChat(ctx context.Context, chatIDs []string, newOwner string) (*response3.ResponseGroupChatTransfer, error) {

	result := &response3.ResponseGroupChatTransfer{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/groupchat/transfer", &object.HashMap{
		"chat_id_list": chatIDs,
		"new_owner":    newOwner,
	}, nil, nil, result)

	return result, err
}

// 查询客户接替结果.
// https://developer.work.weixin.qq.com/document/path/94082
func (comp *Client) GetResignedTransferResult(ctx context.Context, handoverUserID string, takeoverUserID string, cursor string) (*response.ResponseGetTransferedCustomerList, error) {

	result := &response.ResponseGetTransferedCustomerList{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/resigned/transfer_result?", &object.StringMap{
		"handover_userid": handoverUserID,
		"takeover_userid": takeoverUserID,
		"cursor":          cursor,
	}, nil, nil, result)

	return result, err
}
