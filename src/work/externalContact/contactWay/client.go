package contactWay

import (
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	request2 "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/contactWay/request"
	response3 "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/contactWay/response"
)

type Client struct {
	*kernel.BaseClient
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

// 配置客户联系「联系我」方式.
// https://developer.work.weixin.qq.com/document/path/92572
func (comp *Client) Add(options *request2.RequestAddContactWay) (*response3.ResponseAddContactWay, error) {

	result := &response3.ResponseAddContactWay{}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/add_contact_way", options, nil, nil, result)

	return result, err
}

// 获取企业已配置的「联系我」方式
// https://developer.work.weixin.qq.com/document/path/92572
func (comp *Client) Get(configID string) (*response3.ResponseGetContactWay, error) {

	result := &response3.ResponseGetContactWay{}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/get_contact_way", &object.StringMap{
		"config_id": configID,
	}, nil, nil, result)

	return result, err
}

// 获取企业已配置的「联系我」列表
// https://developer.work.weixin.qq.com/document/path/92572
func (comp *Client) List(options *request2.RequestListContactWay) (*response3.ResponseListContactWay, error) {

	result := &response3.ResponseListContactWay{}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/list_contact_way", options, nil, nil, result)

	return result, err
}

// 更新企业已配置的「联系我」方式
// https://developer.work.weixin.qq.com/document/path/92572
func (comp *Client) Update(config *request2.RequestUpdateContactWay) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/update_contact_way", config, nil, nil, result)

	return result, err
}

// 删除企业已配置的「联系我」方式
// https://developer.work.weixin.qq.com/document/path/92572
func (comp *Client) Delete(configID string) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/del_contact_way", &object.StringMap{
		"config_id": configID,
	}, nil, nil, result)

	return result, err
}

// 结束临时会话
// https://developer.work.weixin.qq.com/document/path/92572
func (comp *Client) CloseTempChat(userID string, externalUserID string) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	options := &object.StringMap{
		"userid":          userID,
		"external_userid": externalUserID,
	}
	_, err := comp.HttpPostJson("cgi-bin/externalcontact/close_temp_chat", options, nil, nil, result)

	return result, err
}
