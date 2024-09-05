package customerAcquisition

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	request2 "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/customerAcquisition/request"
	response3 "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/customerAcquisition/response"
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

// 获取获客链接列表
// https://developer.work.weixin.qq.com/document/path/97398#%E8%8E%B7%E5%8F%96%E8%8E%B7%E5%AE%A2%E9%93%BE%E6%8E%A5%E5%88%97%E8%A1%A8
func (comp *Client) ListLink(ctx context.Context, options *request2.RequestListLink) (*response3.ResponseListLink, error) {

	result := &response3.ResponseListLink{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/customer_acquisition/list_link", options, nil, nil, result)

	return result, err
}

// 获取获客链接详情
// https://developer.work.weixin.qq.com/document/path/97398#%E8%8E%B7%E5%8F%96%E8%8E%B7%E5%AE%A2%E9%93%BE%E6%8E%A5%E8%AF%A6%E6%83%85
func (comp *Client) Get(ctx context.Context, linkId string) (*response3.ResponseGetLink, error) {

	result := &response3.ResponseGetLink{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/customer_acquisition/get", &object.StringMap{
		"link_id": linkId,
	}, nil, nil, result)

	return result, err
}

// 创建获客链接
// https://developer.work.weixin.qq.com/document/path/97398#%E5%88%9B%E5%BB%BA%E8%8E%B7%E5%AE%A2%E9%93%BE%E6%8E%A5
func (comp *Client) CreateLink(ctx context.Context, options *request2.RequestCreateLink) (*response3.ResponseCreateLink, error) {

	result := &response3.ResponseCreateLink{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/customer_acquisition/create_link", options, nil, nil, result)

	return result, err
}

// 编辑获客链接
// https://developer.work.weixin.qq.com/document/path/97398#%E7%BC%96%E8%BE%91%E8%8E%B7%E5%AE%A2%E9%93%BE%E6%8E%A5
func (comp *Client) UpdateLink(ctx context.Context, config *request2.RequestUpdateLink) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/customer_acquisition/update_link", config, nil, nil, result)

	return result, err
}

// 删除获客链接
// https://developer.work.weixin.qq.com/document/path/97398#%E5%88%A0%E9%99%A4%E8%8E%B7%E5%AE%A2%E9%93%BE%E6%8E%A5
func (comp *Client) DeleteLink(ctx context.Context, linkId string) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/customer_acquisition/delete_link", &object.StringMap{
		"link_id": linkId,
	}, nil, nil, result)

	return result, err
}
