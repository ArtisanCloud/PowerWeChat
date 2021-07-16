package order

import (
	"github.com/ArtisanCloud/go-libs/object"
	response2 "github.com/ArtisanCloud/power-wechat/src/kernel/response"
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

// Unify order.
// https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_10&index=1

func (comp *Client) Unify(params *object.StringMap, isContract bool) *response2.ResponseWork {

	result := &response2.ResponseWork{}

	endpoint := comp.Wrap("pay/transactions/jsapi")
	comp.Request(endpoint, params, "POST", nil, false, nil, result)

	return result
}

