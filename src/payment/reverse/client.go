package reverse

import (
	"github.com/ArtisanCloud/go-libs/object"
	payment "github.com/ArtisanCloud/power-wechat/src/payment/kernel"
	"github.com/ArtisanCloud/power-wechat/src/payment/reverse/response"
)

type Client struct {
	*payment.BaseClient
}

func NewClient(app *payment.ApplicationPaymentInterface) *Client {
	return &Client{
		payment.NewBaseClient(app),
	}
}

// Reverse order.
// https://pay.weixin.qq.com/wiki/doc/api/micropa

func (comp *Client) Reverse(number string, reverseType string) (interface{}, error) {

	result := &response.ResponseReserve{}

	config := (*comp.App).GetConfig()

	params := &object.HashMap{
		"appid": config.GetString("app_id", ""),
		reverseType: number,
	}

	endpoint := comp.Wrap("secapi/pay/reverse")
	_, err := comp.Request(endpoint, nil, "POST", params, false, nil, result)

	return result, err
}

// Reverse order by out trade number.
// https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_11&index=3

func (comp *Client) ByOutTradeNumber(outTradeNumber string)(interface{}, error) {

	return comp.Reverse(outTradeNumber,"out_trade_no")
}


// Reverse order by transaction_id.
// https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_11&index=3

func (comp *Client) ByTransactionId(transactionID string)(interface{}, error) {

	return comp.Reverse(transactionID,"transaction_id")
}