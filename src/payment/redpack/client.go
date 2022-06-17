package redpack

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	payment "github.com/ArtisanCloud/PowerWeChat/v2/src/payment/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/payment/redpack/request"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/payment/redpack/response"
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

// Query Red Pack.
// https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=13_6&index=5
func (comp *Client) Info(mchBillNO string) (*response.ResponseSendNormal, error) {
	config := (*comp.App).GetConfig()

	result := &response.ResponseSendNormal{}

	params := &object.HashMap{
		"mch_billno": mchBillNO,
		"appid":      config.GetString("app_id", ""),
		"mch_id":     config.GetString("mch_id", ""),
		"bill_type":  "MCHT",
		"nonce_str":  object.QuickRandom(32),
	}

	endpoint := comp.Wrap("mmpaymkttransfers/gethbinfo")
	_, err := comp.SafeRequest(endpoint, params, "POST", &object.HashMap{}, nil, result)

	return result, err
}

// Send Miniprogram Normal redpack.
// https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=18_2&index=2
func (comp *Client) SendMiniProgramNormal(data *request.RequestSendMiniProgramNormal) (*response.ResponseSendMiniProgramNormal, error) {

	result := &response.ResponseSendMiniProgramNormal{}

	config := (*comp.App).GetConfig()

	params, err := object.StructToHashMapWithXML(data)
	if err != nil {
		return nil, err
	}
	base := &object.HashMap{
		"total_num":  config.GetString("mch_id", "1"),
		"wxappid":    config.GetString("app_id", ""),
		"notify_way": "MINI_PROGRAM_JSAPI",
		"mch_id":     config.GetString("mch_id", ""),
	}
	params = object.MergeHashMap(params, base)

	endpoint := comp.Wrap("/mmpaymkttransfers/sendminiprogramhb")
	_, err = comp.SafeRequest(endpoint, params, "POST", &object.HashMap{}, nil, result)

	return result, err
}

// Send Normal redpack.
// https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=13_4&index=3
func (comp *Client) SendNormal(data *request.RequestSendRedPack) (*response.ResponseSendNormal, error) {

	result := &response.ResponseSendNormal{}

	config := (*comp.App).GetConfig()

	params, err := object.StructToHashMapWithXML(data)
	if err != nil {
		return nil, err
	}
	base := &object.HashMap{
		"total_num": 1,
		"wxappid":   config.GetString("app_id", ""),
		"mch_id":    config.GetString("mch_id", ""),
	}
	params = object.MergeHashMap(params, base)

	endpoint := comp.Wrap("/mmpaymkttransfers/sendredpack")
	_, err = comp.SafeRequest(endpoint, params, "POST", &object.HashMap{}, nil, result)

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
	if data.MchID == "" {
		data.MchID = config.GetString("mch_id", "")
	}

	//params, err := object.StructToHashMapWithTag(data,"json")
	params, err := object.StructToHashMap(data)

	endpoint := comp.Wrap("/mmpaymkttransfers/sendgroupredpack")
	_, err = comp.SafeRequest(endpoint, params, "POST", &object.HashMap{}, nil, result)

	return result, err
}
