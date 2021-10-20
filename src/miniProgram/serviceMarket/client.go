package serviceMarket

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/src/miniProgram/serviceMarket/request"
	"github.com/ArtisanCloud/PowerWeChat/src/miniProgram/serviceMarket/response"
)

type Client struct {
	*kernel.BaseClient
}

// 调用服务平台提供的服务
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/service-market/serviceMarket.invokeService.html
func (comp *Client) InvokeService(options *request.RequestServiceMarket) (*response.ResponseServiceMarketInvoceService, error) {

	result := &response.ResponseServiceMarketInvoceService{}

	_, err := comp.HttpPostJson("wxa/servicemarket", options, nil, nil, result)

	return result, err
}
