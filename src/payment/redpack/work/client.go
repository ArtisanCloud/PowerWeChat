package work

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

// Send Work WX redpack.
// https://work.weixin.qq.com/api/doc/90000/90135/90275
func (comp *Client) SendWorkWX(params *request.RequestSendWorkWX) (interface{}, error) {
	config := (*comp.App).GetConfig()

	if params.Wxappid == "" {
		params.Wxappid = config.GetString("app_id", "")
	}

	options, err := object.StructToHashMap(params)

	endpoint := comp.Wrap("/mmpaymkttransfers/sendworkwxredpack")
	result, err := comp.SafeRequest(endpoint, nil, "POST", options, false, nil)

	return result, err
}

// Query Work WX redpack.
// https://work.weixin.qq.com/api/doc/90000/90135/90276
func (comp *Client) QueryWorkWX(params *request.RequestQueryWorkWX) (interface{}, error) {
	config := (*comp.App).GetConfig()

	if params.Appid == "" {
		params.Appid = config.GetString("app_id", "")
	}

	options, err := object.StructToHashMap(params)

	endpoint := comp.Wrap("/mmpaymkttransfers/queryworkwxredpack")
	result, err := comp.SafeRequest(endpoint, nil, "POST", options, false, nil)

	return result, err
}
