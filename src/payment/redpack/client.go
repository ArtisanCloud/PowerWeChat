package redpack

import (
	"github.com/ArtisanCloud/PowerLibs/object"
	payment "github.com/ArtisanCloud/PowerWeChat/src/payment/kernel"
	"github.com/ArtisanCloud/PowerWeChat/src/payment/redpack/request"
	"github.com/ArtisanCloud/PowerWeChat/src/payment/redpack/response"
)

type Client struct {
	*payment.BaseClient
}

func NewClient(app *payment.ApplicationPaymentInterface) *Client {
	return &Client{
		payment.NewBaseClient(app),
	}
}

// Query Red Pack.
// https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=13_6&index=5
func (comp *Client) Info(mchBillNO string) (*response.ResponseSendNormal, error) {
	config := (*comp.App).GetConfig()

	result := &response.ResponseSendNormal{}

	params := &object.StringMap{
		"mch_billno": mchBillNO,
		"appid":      config.GetString("app_id", ""),
		"mch_id":     config.GetString("mch_id", ""),
		"bill_type":  "MCHT",
		"nonce_str":  object.QuickRandom(32),
	}

	endpoint := comp.Wrap("mmpaymkttransfers/gethbinfo")
	_, err := comp.SafeRequest(endpoint, nil, "POST", params, nil, result)

	return result, err
}

// Send Normal redpack.
// https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=13_4&index=3
func (comp *Client) SendNormal(params *request.RequestSendRedPack) (*response.ResponseSendNormal, error) {
	config := (*comp.App).GetConfig()

	result := &response.ResponseSendNormal{}

	externalRequest := (*comp.App).GetExternalRequest()
	clientIP := externalRequest.Host
	if params.ClientIP == "" {
		params.ClientIP = clientIP
	}
	if params.TotalNum == 0 {
		params.TotalNum = 1
	}
	if params.WXAppID == "" {
		params.WXAppID = config.GetString("app_id", "")
	}

	if params.NonceStr == "" {
		params.NonceStr = object.QuickRandom(10)
	}

	options, err := object.StructToStringMap(params)

	endpoint := comp.Wrap("/mmpaymkttransfers/sendredpack")
	_, err = comp.SafeRequest(endpoint, nil, "POST", options, nil, result)

	return result, err
}

// Send Group redpack.
// https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=13_5&index=4
func (comp *Client) SendGroup(params *request.RequestSendGroupRedPack) (*response.ResponseSendGroupRedPack, error) {
	config := (*comp.App).GetConfig()

	result := &response.ResponseSendGroupRedPack{}
	if params.AmtType == "" {
		params.AmtType = "ALL_RAND"
	}
	if params.WXAppID == "" {
		params.WXAppID = config.GetString("app_id", "")
	}
	options, err := object.StructToHashMap(params)

	endpoint := comp.Wrap("/mmpaymkttransfers/sendgroupredpack")
	_, err = comp.SafeRequest(endpoint, nil, "POST", options, nil, result)

	return result, err
}
