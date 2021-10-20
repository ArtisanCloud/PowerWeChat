package redpack

import (
	"github.com/ArtisanCloud/PowerLibs/object"
	payment "github.com/ArtisanCloud/PowerWeChat/src/payment/kernel"
	"github.com/ArtisanCloud/PowerWeChat/src/payment/redpack/request"
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
func (comp *Client) SendMiniProgramNormal(params *request.RequestSendMiniProgramNormal) (interface{}, error) {
	config := (*comp.App).GetConfig()

	if params.TotalNum <= 0 {
		params.TotalNum = 1
	}
	if params.Wxappid == "" {
		params.Wxappid = config.GetString("app_id", "")
	}

	options, err:= object.StructToHashMap(params)

	endpoint := comp.Wrap("/mmpaymkttransfers/sendminiprogramhb")
	result, err := comp.SafeRequest(endpoint, nil, "POST", options, false, nil)

	return result, err
}

