package apply4Sub

import (
	"context"
	"crypto"
	"encoding/base64"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerLibs/v3/security/sign"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/apply4Sub/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/apply4Sub/response"
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

// 特约商户进件
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter11_1_1.shtml
func (comp *Client) ApplyForBusiness(ctx context.Context, params *request.RequestApply) (*response.ResponseApply, error) {

	result := &response.ResponseApply{}

	config := (*comp.App).GetConfig()

	rsaSigner, err := sign.NewRSASigner(crypto.SHA1)
	if err != nil {
		return nil, err
	}
	rsaSigner.RSAEncryptor.PublicKeyPath = config.GetString("rsa_public_key_path", "")
	_, err = rsaSigner.RSAEncryptor.LoadPublicKeyByPath()
	if err != nil {
		return nil, err
	}

	cipherData, err := rsaSigner.RSAEncryptor.Encrypt([]byte(params.ContactInfo.ContactEmail))
	buffer := base64.StdEncoding.EncodeToString(cipherData)
	if err != nil {
		return nil, err
	}
	params.ContactInfo.ContactEmail = string(buffer)
	if err != nil {
		return nil, err
	}

	options, err := object.StructToHashMap(params)
	if err != nil {
		return nil, err
	}

	endpoint := "/v3/applyment4sub/applyment/"
	_, err = comp.Request(ctx, comp.Wrap(endpoint), nil, http.MethodPost, options, false, nil, result)
	return result, err
}
