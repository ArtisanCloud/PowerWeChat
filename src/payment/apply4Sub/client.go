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
func (comp *Client) ApplyForBusiness(ctx context.Context, params *request.RequestApplyForBusiness) (*response.ResponseApplyForBusiness, error) {

	result := &response.ResponseApplyForBusiness{}

	// 获取RSA签名器
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

	// 加密params.ContactInfo.ContactEmail
	cipherData, err := rsaSigner.RSAEncryptor.Encrypt([]byte(params.ContactInfo.ContactEmail))
	buffer := base64.StdEncoding.EncodeToString(cipherData)
	if err != nil {
		return nil, err
	}
	params.ContactInfo.ContactEmail = buffer

	// 加密params.ContactInfo.ContactName
	cipherData, err = rsaSigner.RSAEncryptor.Encrypt([]byte(params.ContactInfo.ContactName))
	buffer = base64.StdEncoding.EncodeToString(cipherData)
	if err != nil {
		return nil, err
	}
	params.ContactInfo.ContactName = buffer

	// 加密params.SubjectInfo.IdentityInfo.IdCardInfo.IdCardNumber
	cipherData, err = rsaSigner.RSAEncryptor.Encrypt([]byte(params.SubjectInfo.IdentityInfo.IdCardInfo.IdCardNumber))
	buffer = base64.StdEncoding.EncodeToString(cipherData)
	if err != nil {
		return nil, err
	}
	params.SubjectInfo.IdentityInfo.IdCardInfo.IdCardNumber = buffer

	// 加密params.SubjectInfo.IdentityInfo.IdCardInfo.IdCardName
	cipherData, err = rsaSigner.RSAEncryptor.Encrypt([]byte(params.SubjectInfo.IdentityInfo.IdCardInfo.IdCardName))
	buffer = base64.StdEncoding.EncodeToString(cipherData)
	if err != nil {
		return nil, err
	}
	params.SubjectInfo.IdentityInfo.IdCardInfo.IdCardName = buffer

	// 加密params.SubjectInfo.IdentityInfo.IdCardInfo.IdCardAddress
	cipherData, err = rsaSigner.RSAEncryptor.Encrypt([]byte(params.SubjectInfo.IdentityInfo.IdCardInfo.IdCardAddress))
	buffer = base64.StdEncoding.EncodeToString(cipherData)
	if err != nil {
		return nil, err
	}
	params.SubjectInfo.IdentityInfo.IdCardInfo.IdCardAddress = buffer

	// 加密params.ContactInfo.mobile_phone
	cipherData, err = rsaSigner.RSAEncryptor.Encrypt([]byte(params.ContactInfo.MobilePhone))
	buffer = base64.StdEncoding.EncodeToString(cipherData)
	if err != nil {
		return nil, err
	}
	params.ContactInfo.MobilePhone = buffer

	// 加密params.BankAccountInfo.AccountName
	cipherData, err = rsaSigner.RSAEncryptor.Encrypt([]byte(params.BankAccountInfo.AccountName))
	buffer = base64.StdEncoding.EncodeToString(cipherData)
	if err != nil {
		return nil, err
	}
	params.BankAccountInfo.AccountName = buffer

	// 加密params.BankAccountInfo.AccountNumber
	cipherData, err = rsaSigner.RSAEncryptor.Encrypt([]byte(params.BankAccountInfo.AccountNumber))
	buffer = base64.StdEncoding.EncodeToString(cipherData)
	if err != nil {
		return nil, err
	}
	params.BankAccountInfo.AccountNumber = buffer

	// 结构体转化
	options, err := object.StructToHashMap(params)
	if err != nil {
		return nil, err
	}

	endpoint := "/v3/applyment4sub/applyment/"
	_, err = comp.Request(ctx, comp.Wrap(endpoint), nil, http.MethodPost, options, false, nil, result)
	return result, err
}

// 通过业务申请编号查询申请状态
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter11_1_2.shtml
func (comp *Client) GetApplyByBusinessCode(ctx context.Context, businessCode string) (*response.ResponseGetApplyForBusiness, error) {

	result := &response.ResponseGetApplyForBusiness{}

	endpoint := "/v3/applyment4sub/applyment/business_code/" + businessCode
	_, err := comp.Request(ctx, endpoint, nil, http.MethodGet, &object.HashMap{}, false, nil, result)
	return result, err
}

// 通过申请单号查询申请状态
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter11_1_2.shtml
func (comp *Client) GetApplyByApplymentId(ctx context.Context, applymentId string) (*response.ResponseGetApplyForBusiness, error) {

	result := &response.ResponseGetApplyForBusiness{}

	endpoint := "/v3/applyment4sub/applyment/applyment_id/" + applymentId
	_, err := comp.Request(ctx, endpoint, nil, http.MethodGet,&object.HashMap{}, false, nil, result)
	return result, err
}
