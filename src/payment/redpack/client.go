package redpack

import (
	payment "github.com/ArtisanCloud/PowerWeChat/src/payment/kernel"
	"github.com/ArtisanCloud/PowerWeChat/src/payment/redpack/request"
	"github.com/ArtisanCloud/go-libs/object"
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
func (comp *Client) Info(mchBillNO string) (interface{}, error) {
	config := (*comp.App).GetConfig()

	params := &object.StringMap{
		"mch_billno": mchBillNO,
		"appid":      config.GetString("app_id", ""),
		"bill_type":  "MCHT",
	}

	endpoint := comp.Wrap("mmpaymkttransfers/gethbinfo")
	result, err := comp.SafeRequest(endpoint, params, "POST", nil, nil, nil)

	return result, err
}

// Send Miniprogram Normal redpack.
// https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=18_2&index=2
func (comp *Client) SendMiniProgramNormal(params *request.SendMiniProgramHB) (interface{}, error) {
	config := (*comp.App).GetConfig()
	if params.TotalNum <= 0 {
		params.TotalNum = 1
	}
	if params.Wxappid != "" {
		params.Wxappid = config.GetString("app_id", "")
	}

	options, err:= object.StructToHashMap(params)

	endpoint := comp.Wrap("/mmpaymkttransfers/sendminiprogramhb")
	result, err := comp.SafeRequest(endpoint, nil, "POST", options, nil, nil)

	return result, err
}

// Send Normal redpack.
// https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=13_4&index=3
func (comp *Client) SendNormal(params *request.SendRedPack) (interface{}, error) {
	config := (*comp.App).GetConfig()

	externalRequest := (*comp.App).GetExternalRequest()
	clientIP := externalRequest.Host
	if params.ClientIp != "" {
		params.ClientIp = clientIP
	}
	if params.TotalNum <= 0 {
		params.TotalNum = 1
	}
	if params.Wxappid != "" {
		params.Wxappid = config.GetString("app_id", "")
	}

	options, err:= object.StructToHashMap(params)

	endpoint := comp.Wrap("/mmpaymkttransfers/sendredpack")
	result, err := comp.SafeRequest(endpoint, nil, "POST", options, nil, nil)

	return result, err
}

// Send Group redpack.
// https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=13_5&index=4
func (comp *Client) SendGroup(params *request.SendGroupRedPack) (interface{}, error) {
	config := (*comp.App).GetConfig()


	if params.AmtType =="" {
		params.AmtType = "ALL_RAND"
	}
	if params.Wxappid != "" {
		params.Wxappid = config.GetString("app_id", "")
	}
	options, err:= object.StructToHashMap(params)

	endpoint := comp.Wrap("/mmpaymkttransfers/sendgroupredpack")
	result, err := comp.SafeRequest(endpoint, nil, "POST", options, nil, nil)

	return result, err
}


// Send Work WX redpack.
// https://work.weixin.qq.com/api/doc/90000/90135/90275
func (comp *Client) SendWorkWX(params *request.SendWorkWXRedPack) (interface{}, error) {
	config := (*comp.App).GetConfig()

	if params.Wxappid != "" {
		params.Wxappid = config.GetString("app_id", "")
	}

	options, err:= object.StructToHashMap(params)

	endpoint := comp.Wrap("/mmpaymkttransfers/sendworkwxredpack")
	result, err := comp.SafeRequest(endpoint, nil, "POST", options, nil, nil)

	return result, err
}


// Query Work WX redpack.
// https://work.weixin.qq.com/api/doc/90000/90135/90276
func (comp *Client) QueryWorkWX(params *request.QueryWorkRedPack) (interface{}, error) {
	config := (*comp.App).GetConfig()

	if params.Appid != "" {
		params.Appid = config.GetString("app_id", "")
	}

	options, err:= object.StructToHashMap(params)

	endpoint := comp.Wrap("/mmpaymkttransfers/queryworkwxredpack")
	result, err := comp.SafeRequest(endpoint, nil, "POST", options, nil, nil)

	return result, err
}