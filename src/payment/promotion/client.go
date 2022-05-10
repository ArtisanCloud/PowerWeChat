package promotion

import (
	payment "github.com/ArtisanCloud/PowerWeChat/src/payment/kernel"
	"github.com/ArtisanCloud/PowerWeChat/src/payment/promotion/request"
	"github.com/ArtisanCloud/PowerWeChat/src/payment/promotion/response"
)

type Client struct {
	*payment.BaseClient
}

func NewClient(app *payment.ApplicationPaymentInterface) *Client {
	return &Client{
		payment.NewBaseClient(app),
	}
}

// 向员工付款
// https://work.weixin.qq.com/api/doc/90000/90135/90278
func (comp *Client) PayTransferToPocket(params *request.RequestPayTransferToPocket) (*response.ResponsePayTransferToPocket, error) {
	result := &response.ResponsePayTransferToPocket{}

	endpoint := comp.Wrap("/mmpaymkttransfers/promotion/paywwsptrans2pocket")
	_, err := comp.SafeRequest(endpoint, nil, "POST", params, false, result)

	return result, err
}

// 查询付款记录
// https://work.weixin.qq.com/api/doc/90000/90135/90279
func (comp *Client) QueryTransferToPocket(params *request.RequestQueryTransferToPocket) (*response.ResponseQueryTransferToPocket, error) {

	result := &response.ResponseQueryTransferToPocket{}

	endpoint := comp.Wrap("/mmpaymkttransfers/promotion/querywwsptrans2pocket")
	_, err := comp.SafeRequest(endpoint, nil, "POST", params, false, result)

	return result, err
}
