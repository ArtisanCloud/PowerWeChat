package transfer

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/transfer/request"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/transfer/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

// 分配在职成员的客户
// https://work.weixin.qq.com/api/doc/90000/90135/92125
func (comp *Client) TransferCustomer(options *request.RequestTransferCustomer) (*response.ResponseTransferCustomer, error) {

	result := &response.ResponseTransferCustomer{}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/transfer_customer", options, nil, nil, result)

	return result, err
}


// 查询客户接替状态
// https://work.weixin.qq.com/api/doc/90000/90135/94088
func (comp *Client) TransferResult(options *request.RequestTransferResult) (*response.ResponseTransferResult, error) {

	result := &response.ResponseTransferResult{}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/transfer_result", options, nil, nil, result)

	return result, err
}



// 获取待分配的离职成员列表
// https://work.weixin.qq.com/api/doc/90000/90135/92125
func (comp *Client) GetUnassignedList(pageID int64, cursor string, pageSize int64) (*response.ResponseResignedGetUnassignedList, error) {

	result := &response.ResponseResignedGetUnassignedList{}

	options := &object.HashMap{
		"page_id":   pageID,
		"cursor":    cursor,
		"page_size": pageSize,
	}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/get_unassigned_list", options, nil, nil, result)

	return result, err
}

// 分配离职成员的客户
// https://work.weixin.qq.com/api/doc/90000/90135/94081
func (comp *Client) ResignedTransferCustomer(options *request.RequestResignedTransferCustomer) (*response.ResponseResignedTransferCustomer, error) {

	result := &response.ResponseResignedTransferCustomer{}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/resigned/transfer_customer", options, nil, nil, result)

	return result, err
}


// 查询客户接替状态
// https://work.weixin.qq.com/api/doc/90000/90135/94082
func (comp *Client) ResignedTransferResult(options *request.RequestResignedTransferCustomer) (*response.ResponseResignedTransferResult, error) {

	result := &response.ResponseResignedTransferResult{}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/resigned/transfer_result", options, nil, nil, result)

	return result, err
}


// 分配离职成员的客户群
// https://work.weixin.qq.com/api/doc/90000/90135/92127
func (comp *Client) GroupChatTransfer(options *request.RequestGroupChatTransfer) (*response.ResponseGroupChatTransfer, error) {

	result := &response.ResponseGroupChatTransfer{}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/groupchat/transfer", options, nil, nil, result)

	return result, err
}

