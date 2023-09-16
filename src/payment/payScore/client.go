package payScore

import (
	"context"
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	payment "github.com/ArtisanCloud/PowerWeChat/v3/src/payment/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/payScore/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/payScore/response"
	"net/http"
)

type Client struct {
	*payment.BaseClient
}

func NewClient(app *payment.ApplicationPaymentInterface) (*Client, error) {
	baseClient, err := payment.NewBaseClient(app)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient,
	}, nil
}

// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_2.shtml
func (comp *Client) ServiceOrder(ctx context.Context, params *request.RequestServiceOrder) (*response.ResponseServiceOrder, error) {

	result := &response.ResponseServiceOrder{}

	options, _ := object.StructToHashMap(params)

	endpoint := comp.Wrap(fmt.Sprintf("/v3/payscore/serviceorder"))
	_, err := comp.Request(ctx, endpoint, nil, http.MethodPost, options, false, nil, result)

	return result, err
}
