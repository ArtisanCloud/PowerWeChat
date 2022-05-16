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
func (comp *Client) SendWorkWX(data *request.RequestSendWorkWX) (*response.ResponseSendWorkWX, error) {
	result := &response.ResponseSendWorkWX{}
	if data.WXAppID == "" {
		config := (*comp.App).GetConfig()
		data.WXAppID = config.GetString("app_id", "")
	}

	params, err := object.StructToHashMap(data)

	endpoint := comp.Wrap("/mmpaymkttransfers/sendworkwxredpack")
	_, err = comp.SafeRequest(endpoint, params, "POST", &object.HashMap{}, false, result)

	return result, err
}

// Query Work WX redpack.
// https://work.weixin.qq.com/api/doc/90000/90135/90276
func (comp *Client) QueryWorkWX(data *request.RequestQueryWorkWX) (*response.ResponseQueryWorkWX, error) {
	result := &response.ResponseQueryWorkWX{}

	if data.Appid == "" {
		config := (*comp.App).GetConfig()
		data.Appid = config.GetString("app_id", "")
	}

	params, err := object.StructToHashMap(data)

	endpoint := comp.Wrap("/mmpaymkttransfers/queryworkwxredpack")
	_, err = comp.SafeRequest(endpoint, params, "POST", &object.HashMap{}, false, result)

	return result, err
}
