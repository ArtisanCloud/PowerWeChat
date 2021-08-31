package refund

import (
	"errors"
	"fmt"
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
		"amount":         amount,
	}

	if transactionID!=""{
		(*body)["transaction_id"]= transactionID
	}else if refundOutNO!=""{
		(*body)["out_refund_no"]= refundOutNO
	}else{
		return nil, errors.New("please given transaction_id or out_refund_no. ")
	}

	body = object.MergeHashMap(body, options)

	endpoint := comp.Wrap("/v3/refund/domestic/refunds")
	_, err := comp.PlainRequest(endpoint, nil, "POST", body, false, nil, result)

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
