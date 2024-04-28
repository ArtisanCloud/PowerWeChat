package merchantService

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	payment "github.com/ArtisanCloud/PowerWeChat/v3/src/payment/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/merchantService/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/merchantService/response"
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

// 查询投诉单列表
// https://pay.weixin.qq.com/docs/merchant/apis/consumer-complaint/complaints/list-complaints-v2.html
func (comp *Client) Complaints(ctx context.Context, params *request.RequestComplaints) (*response.ResponseComplaints, error) {

	result := &response.ResponseComplaints{}

	data, _ := object.StructToStringMap(params)

	endpoint := "/v3/merchant-service/complaints-v2"
	_, err := comp.Request(ctx, endpoint, data, http.MethodGet, &object.HashMap{}, false, nil, result)

	return result, err
}
