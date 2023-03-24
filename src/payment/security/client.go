package security

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	payment "github.com/ArtisanCloud/PowerWeChat/v3/src/payment/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/security/response"
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

// Get Merchant Public Key.
// https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay_yhk.php?chapter=25_7&index=4
func (comp *Client) GetPublicKey(ctx context.Context) (*response.ResponseGetPublicKey, error) {
	config := (*comp.App).GetConfig()

	result := &response.ResponseGetPublicKey{}

	options := &object.HashMap{
		"mch_id": config.GetString("mch_id", ""),
	}

	url := "https://fraud.mch.weixin.qq.com/risk/getpublickey"
	_, err := comp.RequestRawXML(ctx, url, &object.StringMap{}, http.MethodPost, options, nil, result)

	return result, err
}
