package reverse

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	payment "github.com/ArtisanCloud/PowerWeChat/v3/src/payment/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/reverse/response"
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

// Reverse order.
// https://pay.weixin.qq.com/wiki/doc/api/micropa
func (comp *Client) Reverse(ctx context.Context, number string, reverseType string) (*response.ResponseReserve, error) {

	result := &response.ResponseReserve{}

	config := (*comp.App).GetConfig()

	params := &object.HashMap{
		"appid":     config.GetString("app_id", ""),
		reverseType: number,
	}

	endpoint := comp.Wrap("secapi/pay/reverse")
	_, err := comp.SafeRequest(ctx, endpoint, params, http.MethodPost, &object.HashMap{}, nil, result)

	return result, err
}

// Reverse order by out trade number.
// https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_11&index=3
func (comp *Client) ByOutTradeNumber(ctx context.Context, outTradeNumber string) (*response.ResponseReserve, error) {

	return comp.Reverse(ctx, outTradeNumber, "out_trade_no")
}

// Reverse order by transaction_id.
// https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_11&index=3
func (comp *Client) ByTransactionId(ctx context.Context, transactionID string) (*response.ResponseReserve, error) {

	return comp.Reverse(ctx, transactionID, "transaction_id")
}
