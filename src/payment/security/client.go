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

// 获取平台证书
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/wechatpay5_1.shtml
func (comp *Client) GetCertificates(ctx context.Context) (*response.ResponseGetCertificates, error) {

	result := &response.ResponseGetCertificates{}

	endpoint := comp.Wrap("/v3/certificates")
	_, err := comp.Request(ctx, endpoint, nil, http.MethodGet, nil, false, nil, result)

	return result, err
}

// Get RSA Public Key.
// https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay_yhk.php?chapter=25_7&index=4
func (comp *Client) GetRSAPublicKey(ctx context.Context) (*response.ResponseGetPublicKey, error) {
	config := (*comp.App).GetConfig()

	result := &response.ResponseGetPublicKey{}

	options := &object.HashMap{
		"mch_id": config.GetString("mch_id", ""),
	}

	url := "https://fraud.mch.weixin.qq.com/risk/getpublickey"
	_, err := comp.RequestRawXML(ctx, url, &object.StringMap{}, http.MethodPost, options, nil, result)

	return result, err
}
