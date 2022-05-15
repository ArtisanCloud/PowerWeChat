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

// Send Miniprogram Normal redpack.
// https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=18_2&index=2
func (comp *Client) SendMiniProgramNormal(data *request.RequestSendMiniProgramNormal) (*response.ResponseSendMiniProgramNormal, error) {

	result:=&response.ResponseSendMiniProgramNormal{}

	config := (*comp.App).GetConfig()

	if data.TotalNum <= 0 {
		data.TotalNum = 1
	}
	if data.WXAppID == "" {
		data.WXAppID = config.GetString("app_id", "")
	}

	params, err:= object.StructToStringMap(data, "xml")

	endpoint := comp.Wrap("/mmpaymkttransfers/sendminiprogramhb")
	_, err = comp.SafeRequest(endpoint, params, "POST", &object.StringMap{}, false, result)

	return result, err
}

