package order

import (
	"context"
	"errors"
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	payment "github.com/ArtisanCloud/PowerWeChat/v3/src/payment/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/order/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/order/response"
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
func (comp *Client) JSAPITransaction(ctx context.Context, params *request.RequestJSAPIPrepay) (*response.ResponseUnitfy, error) {

	result := &response.ResponseUnitfy{}

	endpoint := "/v3/pay/transactions/jsapi"
	_, err := comp.PayTransaction(ctx, endpoint, params, result)
	return result, err
}

// App pay transaction.
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_2_1.shtml
func (comp *Client) TransactionApp(ctx context.Context, params *request.RequestAppPrepay) (*response.ResponseUnitfy, error) {

	result := &response.ResponseUnitfy{}

	endpoint := "/v3/pay/transactions/app"
	_, err := comp.PayTransaction(ctx, endpoint, params, result)

	return result, err
}

// H5 pay transaction.
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_3_1.shtml
func (comp *Client) TransactionH5(ctx context.Context, params *request.RequestH5Prepay) (*response.ResponseH5URL, error) {

	result := &response.ResponseH5URL{}

	endpoint := "/v3/pay/transactions/h5"
	_, err := comp.PayTransaction(ctx, endpoint, params, result)
	return result, err
}

// Native pay transaction.
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_4_1.shtml
func (comp *Client) TransactionNative(ctx context.Context, params *request.RequestNativePrepay) (*response.ResponseCodeURL, error) {

	result := &response.ResponseCodeURL{}

	endpoint := "/v3/pay/transactions/native"
	_, err := comp.PayTransaction(ctx, endpoint, params, result)

	return result, err
}

// 合单APP下单API
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter5_1_1.shtml
func (comp *Client) TransactionAppCombine(ctx context.Context, params *request.RequestCombinePrepay) (*response.ResponseUnitfy, error) {

	result := &response.ResponseUnitfy{}

	endpoint := "/v3/combine-transactions/app"
	_, err := comp.PayTransaction(ctx, endpoint, params, result)

	return result, err
}

func (comp *Client) PayTransaction(ctx context.Context, entryPoint string, params request.Prepay, result interface{}) (interface{}, error) {
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
	rs, err := comp.Request(ctx, transactionEndpoint, nil, http.MethodPost, options, false, nil, result)
	return rs, err
}

func (comp *Client) QueryByTransactionId(ctx context.Context, number string) (*response.ResponseOrder, error) {
	endpoint := fmt.Sprintf("/v3/pay/transactions/id/%s", number)
	return comp.Query(ctx, &object.HashMap{
		"endpoint": endpoint,
	})
}

func (comp *Client) QueryByOutTradeNumber(ctx context.Context, transactionID string) (*response.ResponseOrder, error) {
	endpoint := fmt.Sprintf("/v3/pay/transactions/out-trade-no/%s", transactionID)
	return comp.Query(ctx, &object.HashMap{
		"endpoint": endpoint,
	})
}

func (comp *Client) Query(ctx context.Context, params *object.HashMap) (*response.ResponseOrder, error) {

	result := &response.ResponseOrder{}

	config := (*comp.App).GetConfig()

	if (*params)["endpoint"] == nil || (*params)["endpoint"].(string) == "" {
		return nil, errors.New("no query endpoint! ")
	}
	endpoint := (*params)["endpoint"].(string)
	endpoint = comp.Wrap(endpoint)
	_, err := comp.Request(ctx, endpoint, &object.StringMap{
		"mchid": config.GetString("mch_id", ""),
	}, http.MethodGet, &object.HashMap{}, false, nil, result)

	return result, err
}

// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_3.shtml
func (comp *Client) Close(ctx context.Context, tradeNo string) (*response2.ResponsePayment, error) {

	result := &response2.ResponsePayment{}

	config := (*comp.App).GetConfig()

	endpoint := comp.Wrap(fmt.Sprintf("/v3/pay/transactions/out-trade-no/%s/close", tradeNo))
	_, err := comp.PlainRequest(ctx, endpoint, nil, http.MethodPost, &object.HashMap{
		"mchid": config.GetString("mch_id", ""),
	}, false, nil, result)

	return result, err
}
