package work

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

// Send Work WX redpack.
// https://work.weixin.qq.com/api/doc/90000/90135/90275
func (comp *Client) SendWorkWX(params *request.RequestSendWorkWX) (*response.ResponseSendWorkWX, error) {
	result := &response.ResponseSendWorkWX{}
	if params.WXAppID == "" {
		config := (*comp.App).GetConfig()
		params.WXAppID = config.GetString("app_id", "")
	}

	options, err := object.StructToHashMap(params)

	endpoint := comp.Wrap("/mmpaymkttransfers/sendworkwxredpack")
	_, err = comp.SafeRequest(endpoint, nil, "POST", options, false, result)

	return result, err
}

// Query Work WX redpack.
// https://work.weixin.qq.com/api/doc/90000/90135/90276
func (comp *Client) QueryWorkWX(params *request.RequestQueryWorkWX) (*response.ResponseQueryWorkWX, error) {
	result := &response.ResponseQueryWorkWX{}

	if params.Appid == "" {
		config := (*comp.App).GetConfig()
		params.Appid = config.GetString("app_id", "")
	}

	options, err := object.StructToHashMap(params)

	endpoint := comp.Wrap("/mmpaymkttransfers/queryworkwxredpack")
	_, err = comp.SafeRequest(endpoint, nil, "POST", options, false, result)

	return result, err
}
