package refund

import (
	"github.com/ArtisanCloud/go-libs/object"
	payment "github.com/ArtisanCloud/power-wechat/src/payment/kernel"
	"github.com/ArtisanCloud/power-wechat/src/payment/refund/response"
)

type Client struct {
	*payment.BaseClient
}

func NewClient(app *payment.ApplicationPaymentInterface) *Client {
	return &Client{
		payment.NewBaseClient(app),
	}
}

// Refund.
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_9.shtml

func (comp *Client) Refund(transactionID string, refundOutNO string, amount *object.HashMap, options *object.HashMap) (*response.ResponseRefund, error) {

	result := &response.ResponseRefund{}

	body := &object.HashMap{
		"transaction_id": transactionID,
		"out_refund_no":  refundOutNO,
		"amount":         amount,
	}
	body = object.MergeHashMap(body, options)

	endpoint := comp.Wrap("/v3/refund/domestic/refunds")
	_, err := comp.Request(endpoint, nil, "POST", body, false, nil, result)

	return result, err
}
