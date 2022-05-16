package promotion

import (
	"github.com/ArtisanCloud/PowerLibs/object"
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
func (comp *Client) PayTransferToPocket(data *request.RequestPayTransferToPocket) (*response.ResponsePayTransferToPocket, error) {
	result := &response.ResponsePayTransferToPocket{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	endpoint := comp.Wrap("/mmpaymkttransfers/promotion/paywwsptrans2pocket")
	_, err = comp.SafeRequest(endpoint, params, "POST", &object.HashMap{}, false, result)

	return result, err
}

// 查询付款记录
// https://work.weixin.qq.com/api/doc/90000/90135/90279
func (comp *Client) QueryTransferToPocket(data *request.RequestQueryTransferToPocket) (*response.ResponseQueryTransferToPocket, error) {

	result := &response.ResponseQueryTransferToPocket{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	endpoint := comp.Wrap("/mmpaymkttransfers/promotion/querywwsptrans2pocket")
	_, err = comp.SafeRequest(endpoint, params, "POST", &object.HashMap{}, false, result)

	return result, err
}
