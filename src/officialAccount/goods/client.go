package goods

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/goods/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/goods/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// 导入商品接口
// https://mp.weixin.qq.com/cgi-bin/announce?action=getannouncement&key=11533749572M9ODP&version=1&lang=zh_CN&platform=2
func (comp *Client) Add(ctx context.Context, data *request.RequestProduct) (*response.ResponseProductAdd, error) {

	result := &response.ResponseProductAdd{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "scan/product/v2/add", data, nil, nil, result)

	return result, err
}

// 更新商品接口
// https://mp.weixin.qq.com/cgi-bin/announce?action=getannouncement&key=11533749572M9ODP&version=1&lang=zh_CN&platform=2
func (comp *Client) Update(ctx context.Context, data *request.RequestProduct) (*response.ResponseProductAdd, error) {

	result := &response.ResponseProductAdd{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "scan/product/v2/add", data, nil, nil, result)

	return result, err
}

// 导入或更新结果查询接口
// https://mp.weixin.qq.com/cgi-bin/announce?action=getannouncement&key=11533749572M9ODP&version=1&lang=zh_CN&platform=2
func (comp *Client) Status(ctx context.Context, ticket string) (*response.ResponseProductStatus, error) {

	result := &response.ResponseProductStatus{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "scan/product/v2/status", &object.HashMap{"status_ticket": ticket}, nil, nil, result)

	return result, err
}

// 单个商品信息查询接口
// https://mp.weixin.qq.com/cgi-bin/announce?action=getannouncement&key=11533749572M9ODP&version=1&lang=zh_CN&platform=2
func (comp *Client) Get(ctx context.Context, data *request.RequestProductGet) (*response.ResponseProductGet, error) {

	result := &response.ResponseProductGet{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "scan/product/v2/getinfo", data, nil, nil, result)

	return result, err
}

// 全量商品信息查询接口
// https://mp.weixin.qq.com/cgi-bin/announce?action=getannouncement&key=11533749572M9ODP&version=1&lang=zh_CN&platform=2
func (comp *Client) List(ctx context.Context, context string, page int, size int) (*response.ResponseProductGet, error) {

	result := &response.ResponseProductGet{}

	params := &object.HashMap{
		"page_context": context,
		"page_num":     page,
		"page_size":    size,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "scan/product/v2/getinfobypage", params, nil, nil, result)

	return result, err
}
