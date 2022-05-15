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
	_, err := comp.SafeRequest(endpoint, params, "POST", &object.StringMap{}, nil, result)

	return result, err
}

// Send Normal redpack.
// https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=13_4&index=3
func (comp *Client) SendNormal(data *request.RequestSendRedPack) (*response.ResponseSendNormal, error) {
	config := (*comp.App).GetConfig()

	result := &response.ResponseSendNormal{}

	externalRequest := (*comp.App).GetExternalRequest()
	clientIP := externalRequest.Host
	if data.ClientIP == "" {
		data.ClientIP = clientIP
	}
	if data.TotalNum == 0 {
		data.TotalNum = 1
	}
	if data.WXAppID == "" {
		data.WXAppID = config.GetString("app_id", "")
	}

	if data.NonceStr == "" {
		data.NonceStr = object.QuickRandom(10)
	}

	params, err := object.StructToStringMap(data, "xml")

	endpoint := comp.Wrap("/mmpaymkttransfers/sendredpack")
	_, err = comp.SafeRequest(endpoint, params, "POST", &object.StringMap{}, nil, result)

	return result, err
}

// Send Group redpack.
// https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=13_5&index=4
func (comp *Client) SendGroup(data *request.RequestSendGroupRedPack) (*response.ResponseSendGroupRedPack, error) {
	config := (*comp.App).GetConfig()

	result := &response.ResponseSendGroupRedPack{}
	if data.AmtType == "" {
		data.AmtType = "ALL_RAND"
	}
	if data.WXAppID == "" {
		data.WXAppID = config.GetString("app_id", "")
	}
	params, err := object.StructToStringMap(data, "xml")

	endpoint := comp.Wrap("/mmpaymkttransfers/sendgroupredpack")
	_, err = comp.SafeRequest(endpoint, params, "POST", &object.StringMap{}, nil, result)

	return result, err
}
