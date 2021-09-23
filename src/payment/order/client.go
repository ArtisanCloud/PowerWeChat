package order

import (
	"errors"
	"fmt"
	response2 "github.com/ArtisanCloud/go-libs/http/response"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
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

// JSAPI pay transaction. 小程序支付
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_1.shtml
func (comp *Client) JSAPITransaction(params *power.HashMap) (*response.ResponseUnitfy, error) {

	result := &response.ResponseUnitfy{}

	endpoint := "/v3/pay/transactions/jsapi"
	_, err := comp.PayTransaction(endpoint, params, result)
	return result, err
}

// App pay transaction.
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_2_1.shtml
func (comp *Client) TransactionApp(params *power.HashMap) (*response.ResponseUnitfy, error) {

	result := &response.ResponseUnitfy{}

	endpoint := "/v3/pay/transactions/app"
	_, err := comp.PayTransaction(endpoint, params, result)

	return result, err
}

// H5 pay transaction.
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_3_1.shtml
func (comp *Client) TransactionH5(params *power.HashMap) (*response.ResponseH5URL, error) {

	result := &response.ResponseH5URL{}

	endpoint := "/v3/pay/transactions/h5"
	_, err := comp.PayTransaction(endpoint, params, result)
	return result, err
}

// Native pay transaction.
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_4_1.shtml
func (comp *Client) TransactionNative(params *power.HashMap) (*response.ResponseCodeURL, error) {

	result := &response.ResponseCodeURL{}

	endpoint := "/v3/pay/transactions/native"
	_, err := comp.PayTransaction(endpoint, params, result)

	return result, err
}

// 合单APP下单API
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter5_1_1.shtml
func (comp *Client) TransactionAppCombine(params *power.HashMap) (*response.ResponseUnitfy, error) {

	result := &response.ResponseUnitfy{}

	endpoint := "/v3/combine-transactions/app"
	_, err := comp.PayTransaction(endpoint, params, result)

	return result, err
}

func (comp *Client) PayTransaction(entryPoint string, params *power.HashMap, result interface{}) (interface{}, error) {
	config := (*comp.App).GetConfig()
	(*params)["appid"] = config.GetString("app_id", "")
	if (*params)["notify_url"] == nil || (*params)["notify_url"] == "" {
		(*params)["notify_url"] = config.GetString("notify_url", "")
	}

	transactionEndpoint := comp.Wrap(entryPoint)
	rs, err := comp.Request(transactionEndpoint, nil, "POST", params.ToHashMap(), false, nil, result)
	return rs, err
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
