package transfer

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/transfer/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/transfer/response"
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

// 分配在职成员的客户
// https://developer.work.weixin.qq.com/document/path/92125
func (comp *Client) TransferCustomer(ctx context.Context, options *request.RequestTransferCustomer) (*response.ResponseTransferCustomer, error) {

	result := &response.ResponseTransferCustomer{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/transfer_customer", options, nil, nil, result)

	return result, err
}

// 查询客户接替状态
// https://developer.work.weixin.qq.com/document/path/94088
func (comp *Client) TransferResult(ctx context.Context, options *request.RequestTransferResult) (*response.ResponseTransferResult, error) {

	result := &response.ResponseTransferResult{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/transfer_result", options, nil, nil, result)

	return result, err
}

// 获取待分配的离职成员列表
// https://developer.work.weixin.qq.com/document/path/92125
func (comp *Client) GetUnassignedList(ctx context.Context, pageID int, cursor string, pageSize int) (*response.ResponseResignedGetUnassignedList, error) {

	result := &response.ResponseResignedGetUnassignedList{}

	options := &object.HashMap{
		"page_id":   pageID,
		"cursor":    cursor,
		"page_size": pageSize,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/get_unassigned_list", options, nil, nil, result)

	return result, err
}

// 分配离职成员的客户
// https://developer.work.weixin.qq.com/document/path/94081
func (comp *Client) ResignedTransferCustomer(ctx context.Context, options *request.RequestResignedTransferCustomer) (*response.ResponseResignedTransferCustomer, error) {

	result := &response.ResponseResignedTransferCustomer{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/resigned/transfer_customer", options, nil, nil, result)

	return result, err
}

// 查询客户接替状态
// https://developer.work.weixin.qq.com/document/path/94082
func (comp *Client) ResignedTransferResult(ctx context.Context, options *request.RequestResignedTransferResult) (*response.ResponseResignedTransferResult, error) {

	result := &response.ResponseResignedTransferResult{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/resigned/transfer_result", options, nil, nil, result)

	return result, err
}

// 分配离职成员的客户群
// https://developer.work.weixin.qq.com/document/path/92127
func (comp *Client) GroupChatTransfer(ctx context.Context, options *request.RequestGroupChatTransfer) (*response.ResponseGroupChatTransfer, error) {

	result := &response.ResponseGroupChatTransfer{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/groupchat/transfer", options, nil, nil, result)

	return result, err
}
