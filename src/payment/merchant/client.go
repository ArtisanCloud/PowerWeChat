package merchant

import (
	"context"
	"errors"
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	payment "github.com/ArtisanCloud/PowerWeChat/v3/src/payment/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/merchant/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/merchant/response"

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

// 商户系统先调用该接口在微信支付服务后台生成预支付交易单，返回正确的预支付交易会话标识后再按Native、JSAPI、APP等不同场景生成交易串调起支付。
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_1_1.shtml
func (comp *Client) JSAPITransaction(ctx context.Context, params *request.RequestJSAPIPrepay) (*response.ResponseUnitfy, error) {

	result := &response.ResponseUnitfy{}

	endpoint := "/v3/pay/partner/transactions/jsapi"
	_, err := comp.PayTransaction(ctx, endpoint, params, result)
	return result, err
}

// 商户系统先调用该接口在微信支付服务后台生成预支付交易单，返回正确的预支付交易会话标识后再按Native、JSAPI、APP等不同场景生成交易串调起支付。
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_2_1.shtml
func (comp *Client) TransactionApp(ctx context.Context, params *request.RequestAppPrepay) (*response.ResponseUnitfy, error) {

	result := &response.ResponseUnitfy{}

	endpoint := "/v3/pay/partner/transactions/app"
	_, err := comp.PayTransaction(ctx, endpoint, params, result)

	return result, err
}

// 商户系统先调用该接口在微信支付服务后台生成预支付交易单，返回正确的预支付交易会话标识后再按Native、JSAPI、APP等不同场景生成交易串调起支付。
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_3_1.shtml
func (comp *Client) TransactionH5(ctx context.Context, params *request.RequestH5Prepay) (*response.ResponseH5URL, error) {

	result := &response.ResponseH5URL{}

	endpoint := "/v3/pay/partner/transactions/h5"
	_, err := comp.PayTransaction(ctx, endpoint, params, result)
	return result, err
}

// 商户Native下单接口，微信后台系统返回链接参数code_url，商户后台系统将code_url值生成二维码图片，用户使用微信客户端扫码后发起支付。
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_4_1.shtml
func (comp *Client) TransactionNative(ctx context.Context, params *request.RequestNativePrepay) (*response.ResponseCodeURL, error) {

	result := &response.ResponseCodeURL{}

	endpoint := "/v3/pay/partner/transactions/native"
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

	if params.GetSpAppId() == "" {
		appID := config.GetString("app_id", "")
		params.SetSpAppId(appID)
	}

	if params.GetSubAppId() == "" {
		appID := config.GetString("sub_appid", "")
		params.SetSubAppId(appID)
	}
	if params.GetNotifyUrl() == "" {
		url := config.GetString("notify_url", "")
		params.SetNotifyUrl(url)
	}

	if params.GetSubAppId() == "" {
		appID := config.GetString("sub_appid", "")
		params.SetSubAppId(appID)
	}

	if params.GetSpMchId() == "" {
		mchID := config.GetString("mch_id", "")
		params.SetSpMchId(mchID)
	}
	if params.GetSubMchId() == "" {
		mchID := config.GetString("sub_mch_id", "")
		params.SetSubMchId(mchID)
	}

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
	endpoint := fmt.Sprintf("/v3/pay/partner/transactions/id/%s", number)
	return comp.Query(ctx, &object.HashMap{
		"endpoint": endpoint,
	})
}

func (comp *Client) QueryByOutTradeNumber(ctx context.Context, transactionID string) (*response.ResponseOrder, error) {
	endpoint := fmt.Sprintf("/v3/pay/partner/transactions/out-trade-no/%s", transactionID)
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
		"sp_mchid":  config.GetString("mch_id", ""),
		"sub_mchid": config.GetString("sub_mch_id", ""),
	}, http.MethodGet, nil, false, nil, result)

	return result, err
}

// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_3.shtml
func (comp *Client) Close(ctx context.Context, tradeNo string) (*http.Response, error) {

	config := (*comp.App).GetConfig()

	endpoint := comp.Wrap(fmt.Sprintf("/v3/pay/partner/transactions/out-trade-no/%s/close", tradeNo))
	rs, err := comp.PlainRequest(ctx, endpoint, nil, http.MethodPost, &object.HashMap{
		"sp_mchid":  config.GetString("mch_id", ""),
		"sub_mchid": config.GetString("sub_mch_id", ""),
	}, false, nil, nil)

	return rs, err
}
