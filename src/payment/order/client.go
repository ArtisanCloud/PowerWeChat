package order

import (
	"errors"
	"fmt"
	response2 "github.com/ArtisanCloud/PowerLibs/v2/http/response"
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	payment "github.com/ArtisanCloud/PowerWeChat/v2/src/payment/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/payment/order/request"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/payment/order/response"
	"net/http"
)

type Client struct {
	*payment.BaseClient
}

func NewClient(app *payment.ApplicationPaymentInterface) (*Client, error) {
	baseClient, err := payment.NewBaseClient(app)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient,
	}, nil
}

// JSAPI pay transaction. 小程序支付
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_1.shtml
func (comp *Client) JSAPITransaction(params *request.RequestJSAPIPrepay) (*response.ResponseUnitfy, error) {

	result := &response.ResponseUnitfy{}

	endpoint := "/v3/pay/transactions/jsapi"
	_, err := comp.PayTransaction(endpoint, params, result)
	return result, err
}

// App pay transaction.
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_2_1.shtml
func (comp *Client) TransactionApp(params *request.RequestAppPrepay) (*response.ResponseUnitfy, error) {

	result := &response.ResponseUnitfy{}

	endpoint := "/v3/pay/transactions/app"
	_, err := comp.PayTransaction(endpoint, params, result)

	return result, err
}

// H5 pay transaction.
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_3_1.shtml
func (comp *Client) TransactionH5(params *request.RequestH5Prepay) (*response.ResponseH5URL, error) {

	result := &response.ResponseH5URL{}

	endpoint := "/v3/pay/transactions/h5"
	_, err := comp.PayTransaction(endpoint, params, result)
	return result, err
}

// Native pay transaction.
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_4_1.shtml
func (comp *Client) TransactionNative(params *request.RequestNativePrepay) (*response.ResponseCodeURL, error) {

	result := &response.ResponseCodeURL{}

	endpoint := "/v3/pay/transactions/native"
	_, err := comp.PayTransaction(endpoint, params, result)

	return result, err
}

// 合单APP下单API
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter5_1_1.shtml
func (comp *Client) TransactionAppCombine(params *request.RequestCombinePrepay) (*response.ResponseUnitfy, error) {

	result := &response.ResponseUnitfy{}

	endpoint := "/v3/combine-transactions/app"
	_, err := comp.PayTransaction(endpoint, params, result)

	return result, err
}

func (comp *Client) PayTransaction(entryPoint string, params request.Prepay, result interface{}) (interface{}, error) {
	config := (*comp.App).GetConfig()

	if params.GetAppID() == "" {
		appID := config.GetString("app_id", "")
		params.SetAppID(appID)
	}
	if params.GetNotifyUrl() == "" {
		url := config.GetString("notify_url", "")
		params.SetNotifyUrl(url)
	}

	mchID := config.GetString("mch_id", "")
	params.SetMchID(mchID)

	//options, err := object.StructToHashMapWithTag(params,"json")
	options, err := object.StructToHashMap(params)
	if err != nil {
		return nil, err
	}

	transactionEndpoint := comp.Wrap(entryPoint)
	rs, err := comp.Request(transactionEndpoint, nil, "POST", options, false, nil, result)
	return rs, err
}

func (comp *Client) QueryByTransactionId(number string) (*response.ResponseOrder, error) {
	endpoint := fmt.Sprintf("/v3/pay/transactions/id/%s", number)
	return comp.Query(&object.HashMap{
		"endpoint": endpoint,
	})
}

func (comp *Client) QueryByOutTradeNumber(transactionID string) (*response.ResponseOrder, error) {
	endpoint := fmt.Sprintf("/v3/pay/transactions/out-trade-no/%s", transactionID)
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
