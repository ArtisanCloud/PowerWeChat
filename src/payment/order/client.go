package order

import (
	"fmt"
	"github.com/ArtisanCloud/go-libs/object"
	payment "github.com/ArtisanCloud/power-wechat/src/payment/kernel"
	"github.com/ArtisanCloud/power-wechat/src/payment/order/response"
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
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_2.shtml

func (comp *Client) Unify(params *object.HashMap, isContract bool) *response.ResponseUnitfy {

	result := &response.ResponseUnitfy{}

	config := (*comp.App).GetConfig()
	(*params)["appid"] = config.GetString("app_id", "")
	if (*params)["notify_url"] == nil || (*params)["notify_url"] == "" {
		(*params)["notify_url"] = config.GetString("notify_url", "")
	}

	endpoint := comp.Wrap("pay/transactions/jsapi")
	comp.Request(endpoint, nil, "POST", params, false, nil, result)

	return result
}
func (comp *Client) QueryByOutTradeNumber(number string) *response.ResponseOrder {
	endpoint := fmt.Sprintf("pay/transactions/out-trade-no/transaction_id/%s", number)
	return comp.Query(endpoint)
}

func (comp *Client) QueryByTransactionId(transactionId string) *response.ResponseOrder {
	endpoint := fmt.Sprintf("pay/transactions/id/%s", transactionId)
	return comp.Query(endpoint)
}

func (comp *Client) Query(endpoint string) *response.ResponseOrder {

	result := &response.ResponseOrder{}

	config := (*comp.App).GetConfig()
	endpoint = comp.Wrap(endpoint)
	comp.Request(endpoint, &object.StringMap{
		"appid": config.GetString("app_id", ""),
	}, "GET", nil, false, nil, result)

	return result
}

func (comp *Client) Close(tradeNo string) *response.ResponseHeaderCloseOrdr {

	result := &response.ResponseHeaderCloseOrdr{}

	config := (*comp.App).GetConfig()

	endpoint := comp.Wrap(fmt.Sprintf("pay/transactions/out-trade-no/%s/close", tradeNo))
	comp.Request(endpoint, &object.StringMap{
		"appid":        config.GetString("app_id", ""),
		"out_trade_no": tradeNo,
	}, "POST", nil, false, result, nil)

	return result
}
