package serviceMarket

import (
	"context"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/serviceMarket/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/serviceMarket/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// 调用服务平台提供的服务
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/service-market/serviceMarket.invokeService.html
func (comp *Client) InvokeService(ctx context.Context, options *request.RequestServiceMarket) (*response.ResponseServiceMarketInvoceService, error) {

	result := &response.ResponseServiceMarketInvoceService{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/servicemarket", options, nil, nil, result)

	return result, err
}
