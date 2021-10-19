package serviceMarket

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/powerwechat/src/kernel"
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/miniProgram/serviceMarket/response"
)

type Client struct {
	*kernel.BaseClient
}

// 调用服务平台提供的服务
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/service-market/serviceMarket.invokeService.html
func (comp *Client) InvokeService(service string, api string, serviceData *power.HashMap, clientMsgID string) (*response.ResponseServiceMarketInvoceService, error) {

	result := &response.ResponseServiceMarketInvoceService{}

	data := &object.HashMap{
		"service":       service,
		"api":           api,
		"data":          serviceData,
		"client_msg_id": clientMsgID,
	}

	_, err := comp.HttpPostJson("wxa/servicemarket", data, nil, nil, result)

	return result, err
}
