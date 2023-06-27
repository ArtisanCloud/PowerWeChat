package tax

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/tax/response"
	"net/http"

	payment "github.com/ArtisanCloud/PowerWeChat/v3/src/payment/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/tax/request"
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

// 创建电子发票卡券模板
// https://pay.weixin.qq.com/docs/merchant/apis/fapiao/fapiao-card-template/create-fapiao-card-template.html
func (comp *Client) ApplyForCardTemplate(ctx context.Context, params *request.RequestApplyForCardTemplate) (*response.ResponseApplyForCardTemplate, error) {

	result := &response.ResponseApplyForCardTemplate{}

	options, _ := object.StructToHashMap(params)

	endpoint := "/v3/new-tax-control-fapiao/card-template"

	_, err := comp.Request(ctx, comp.Wrap(endpoint), nil, http.MethodPost, options, false, nil, result)
	return result, err
}
