package promotion

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	payment "github.com/ArtisanCloud/power-wechat/src/payment/kernel"
)

type Client struct {
	*payment.BaseClient
}

func NewClient(app *payment.ApplicationPaymentInterface) *Client {
	return &Client{
		payment.NewBaseClient(app),
	}
}

// 向员工付款
// https://work.weixin.qq.com/api/doc/90000/90135/90278
func (comp *Client) PayTransferToPocket(params *power.HashMap) (interface{}, error) {

	config := (*comp.App).GetConfig()

	base := &object.HashMap{
		"wxappid":  config.GetString("app_id", ""),
	}

	options := object.MergeHashMap(base, params.ToHashMap())

	endpoint := comp.Wrap("/mmpaymkttransfers/promotion/paywwsptrans2pocket")
	result, err := comp.SafeRequest(endpoint, nil, "POST", options, false, nil)

	return result, err
}


// 查询付款记录
// https://work.weixin.qq.com/api/doc/90000/90135/90279
func (comp *Client) QueryTransferToPocket(params *power.HashMap) (interface{}, error) {

	config := (*comp.App).GetConfig()

	base := &object.HashMap{
		"wxappid":  config.GetString("app_id", ""),
	}

	options := object.MergeHashMap(base, params.ToHashMap())

	endpoint := comp.Wrap("/mmpaymkttransfers/promotion/querywwsptrans2pocket")
	result, err := comp.SafeRequest(endpoint, nil, "POST", options, false, nil)

	return result, err
}