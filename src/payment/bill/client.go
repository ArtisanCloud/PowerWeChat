package bill

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/bill/response"
	payment "github.com/ArtisanCloud/PowerWeChat/v3/src/payment/kernel"
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

// Get Trade Bill.
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_6.shtml
func (comp *Client) GetTradeBill(ctx context.Context, date string, billType string, tarType string) (*response.ResponseBillGet, error) {

	result := &response.ResponseBillGet{}

	params := &object.StringMap{
		"bill_date": date,
		"bill_type": billType,
		"tar_type":  tarType,
	}

	endpoint := comp.Wrap("/v3/bill/tradebill")
	_, err := comp.Request(ctx, endpoint, params, http.MethodPost, &object.HashMap{}, false, nil, result)

	return result, err
}

// Get Flow Bill.
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_7.shtml
func (comp *Client) GetFlowBill(ctx context.Context, date string, accountType string, tarType string) (*response.ResponseBillGet, error) {

	result := &response.ResponseBillGet{}

	params := &object.StringMap{
		"bill_date":    date,
		"account_type": accountType,
		"tar_type":     tarType,
	}

	endpoint := comp.Wrap("/v3/bill/fundflowbill")
	_, err := comp.Request(ctx, endpoint, params, http.MethodPost, &object.HashMap{}, false, nil, result)

	return result, err
}

// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_8.shtml
func (comp *Client) DownloadBill(ctx context.Context, requestDownload *power.RequestDownload, filePath string) (int64, error) {
	return comp.StreamDownload(ctx, requestDownload, filePath)
}
