package order

import (
	"errors"
	"fmt"
	response2 "github.com/ArtisanCloud/go-libs/http/response"
	"github.com/ArtisanCloud/go-libs/object"
	payment "github.com/ArtisanCloud/power-wechat/src/payment/kernel"
	"github.com/ArtisanCloud/power-wechat/src/payment/order/response"
	"net/http"
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

func (comp *Client) Unify(params *object.HashMap, isContract bool) (*response.ResponseUnitfy, error) {

	result := &response.ResponseUnitfy{}

	config := (*comp.App).GetConfig()
	(*params)["appid"] = config.GetString("app_id", "")
	if (*params)["notify_url"] == nil || (*params)["notify_url"] == "" {
		(*params)["notify_url"] = config.GetString("notify_url", "")
	}

	endpoint := comp.Wrap("/v3/pay/transactions/jsapi")
	_, err := comp.Request(endpoint, nil, "POST", params, false, nil, result)

	return result, err
}

func (comp *Client) QueryByTransactionId(number string) (*response.ResponseOrder, error) {
	endpoint := fmt.Sprintf("/v3/pay/transactions/id/%s", number)
	return comp.Query(&object.HashMap{
		"endpoint": endpoint,
	})
}

func (comp *Client) QueryByOutTradeNumber(transactionId string) (*response.ResponseOrder, error) {
	endpoint := fmt.Sprintf("/v3/pay/transactions/out-trade-no/%s", transactionId)
	return comp.Query(&object.HashMap{
		"endpoint": endpoint,
	})
}

func (comp *Client) Query(params *object.HashMap) (*response.ResponseOrder, error) {

	result := &response.ResponseOrder{}

	config := (*comp.App).GetConfig()

	if (*params)["endpoint"] == nil || (*params)["endpoint"].(string) == "" {
		return nil, errors.New("no query endpoint! ")
	}
	endpoint := (*params)["endpoint"].(string)
	endpoint = comp.Wrap(endpoint)
	_, err := comp.Request(endpoint, &object.StringMap{
		"mchid": config.GetString("mch_id", ""),
	}, "GET", nil, false, nil, result)

	return result, err
}

// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_3.shtml
func (comp *Client) Close(tradeNo string) (*http.Response, error) {

	config := (*comp.App).GetConfig()

	endpoint := comp.Wrap(fmt.Sprintf("/v3/pay/transactions/out-trade-no/%s/close", tradeNo))
	rs, err := comp.PlainRequest(endpoint, nil, "POST", &object.HashMap{
		"mchid": config.GetString("mch_id", ""),
	}, false, nil, nil)

	var httpResponse *http.Response = nil
	if err != nil {
		httpResponse = rs.(*response2.HttpResponse).Response
	}

	return httpResponse, err
}
