package work

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	payment "github.com/ArtisanCloud/PowerWeChat/src/payment/kernel"
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
func (comp *Client) SendWorkWX(params *power.HashMap) (interface{}, error) {
	config := (*comp.App).GetConfig()

	base := &object.HashMap{
		"wxappid":  config.GetString("app_id", ""),
	}

	options := object.MergeHashMap(base, params.ToHashMap())

	endpoint := comp.Wrap("/mmpaymkttransfers/sendworkwxredpack")
	result, err := comp.SafeRequest(endpoint, nil, "POST", options, false, nil)

	return result, err
}


// Query Work WX redpack.
// https://work.weixin.qq.com/api/doc/90000/90135/90276
func (comp *Client) QueryWorkWX(params *power.HashMap) (interface{}, error) {
	config := (*comp.App).GetConfig()

	base := &object.HashMap{
		"wxappid":  config.GetString("app_id", ""),
	}

	options := object.MergeHashMap(base, params.ToHashMap())

	endpoint := comp.Wrap("/mmpaymkttransfers/queryworkwxredpack")
	result, err := comp.SafeRequest(endpoint, nil, "POST", options, false, nil)

	return result, err
}