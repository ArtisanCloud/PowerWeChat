package refund

import (
	"fmt"
	"github.com/ArtisanCloud/go-libs/object"
	payment "github.com/ArtisanCloud/powerwechat/src/payment/kernel"
	"github.com/ArtisanCloud/powerwechat/src/payment/refund/request"
	"github.com/ArtisanCloud/powerwechat/src/payment/refund/response"
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

func (comp *Client) Refund(options *request.RequestRefund) (*response.ResponseRefund, error) {

	result := &response.ResponseRefund{}

	body, err := object.StructToHashMap(options)
	if err != nil {
		return nil, err
	}

	endpoint := comp.Wrap("/v3/refund/domestic/refunds")
	_, err = comp.PlainRequest(endpoint, nil, "POST", body, false, nil, result)

	return result, err
}

// Query Refund.
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_10.shtml

func (comp *Client) Query(refundOutNO string) (*response.ResponseRefund, error) {

	result := &response.ResponseRefund{}

	endpoint := fmt.Sprintf("/v3/refund/domestic/refunds/%s", refundOutNO)
	endpoint = comp.Wrap(endpoint)
	_, err := comp.Request(endpoint, nil, "GET", nil, false, nil, result)

	return result, err
}
