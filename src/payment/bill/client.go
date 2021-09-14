package bill

import (
	response2 "github.com/ArtisanCloud/go-libs/http/response"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/payment/bill/response"
	payment "github.com/ArtisanCloud/power-wechat/src/payment/kernel"
)

type Client struct {
	*payment.BaseClient
}

func NewClient(app *payment.ApplicationPaymentInterface) *Client {
	return &Client{
		payment.NewBaseClient(app),
	}
}

// Get Trade Bill.
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_6.shtml
func (comp *Client) GetTradeBill(date string, billType string, tarType string) (*response.ResponseBillGet, error) {

	result := &response.ResponseBillGet{}

	params := &object.StringMap{
		"bill_date": date,
		"bill_type": billType,
		"tar_type":  tarType,
	}

	endpoint := comp.Wrap("/v3/bill/tradebill")
	_, err := comp.Request(endpoint, params, "GET", nil, false, nil, result)

	return result, err
}

// Get Flow Bill.
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_7.shtml
func (comp *Client) GetFlowBill(date string, accountType string, tarType string) (*response.ResponseBillGet, error) {

	result := &response.ResponseBillGet{}

	params := &object.StringMap{
		"bill_date":    date,
		"account_type": accountType,
		"tar_type":     tarType,
	}

	endpoint := comp.Wrap("/v3/bill/fundflowbill")
	_, err := comp.Request(endpoint, params, "GET", nil, false, nil, result)

	return result, err
}

// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_8.shtml
func (comp *Client) DownloadBill(requestDownload *response2.RequestDownload, filePath string) (int64, error) {
	return comp.StreamDownload(requestDownload, filePath)
}
