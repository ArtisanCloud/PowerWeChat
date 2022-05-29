package moment

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/work/externalContact/moment/request"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/work/externalContact/moment/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

// 获取企业全部的发表列表
// https://work.weixin.qq.com/api/doc/90000/90135/93333
func (comp *Client) GetMomentList(params *request.RequestGetMomentList) (*response.ResponseGetMomentList, error) {

	result := &response.ResponseGetMomentList{}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/get_moment_list", params, nil, nil, result)

	return result, err
}

// 获取客户朋友圈企业发表的列表
// https://work.weixin.qq.com/api/doc/90000/90135/93333
func (comp *Client) GetMomentTask(momentID string, cursor string, limit int) (*response.ResponseMomentGetMomentTask, error) {

	result := &response.ResponseMomentGetMomentTask{}

	options := &object.HashMap{
		"moment_id": momentID,
		"cursor":    cursor,
		"limit":     limit,
	}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/get_moment_task", options, nil, nil, result)

	return result, err
}

// 获取客户朋友圈发表时选择的可见范围
// https://work.weixin.qq.com/api/doc/90000/90135/93333
func (comp *Client) GetMomentCustomerList(momentID string, userID string, cursor string, limit int) (*response.ResponseMomentGetMomentCustomerList, error) {

	result := &response.ResponseMomentGetMomentCustomerList{}

	options := &object.HashMap{
		"moment_id": momentID,
		"userid":    userID,
		"cursor":    cursor,
		"limit":     limit,
	}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/get_moment_customer_list", options, nil, nil, result)

	return result, err
}

// 获取客户朋友圈发表后的可见客户列表
// https://work.weixin.qq.com/api/doc/90000/90135/93333
func (comp *Client) GetMomentSendResult(momentID string, userID string, cursor string, limit int) (*response.ResponseMomentGetMomentSendResult, error) {

	result := &response.ResponseMomentGetMomentSendResult{}

	options := &object.HashMap{
		"moment_id": momentID,
		"userid":    userID,
		"cursor":    cursor,
		"limit":     limit,
	}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/get_moment_send_result", options, nil, nil, result)

	return result, err
}

// 获取客户朋友圈的互动数据
// https://work.weixin.qq.com/api/doc/90000/90135/93333
func (comp *Client) GetMomentComments(momentID string, userID string) (*response.ResponseMomentGetMomentComments, error) {

	result := &response.ResponseMomentGetMomentComments{}

	options := &object.HashMap{
		"moment_id": momentID,
		"userid":    userID,
	}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/get_moment_comments", options, nil, nil, result)

	return result, err
}
