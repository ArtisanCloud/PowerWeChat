package transfer

import (
	"context"
	"crypto"
	"encoding/base64"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerLibs/v3/security/sign"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	payment "github.com/ArtisanCloud/PowerWeChat/v3/src/payment/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/transfer/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/transfer/response"
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

// Query MerchantPay to balance.
// https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_3
func (comp *Client) QueryBalanceOrder(ctx context.Context, partnerTradeNo string) (*response.ResponseGetTransferInfo, error) {
	config := (*comp.App).GetConfig()

	result := &response.ResponseGetTransferInfo{}

	params := &object.HashMap{
		"appid":            config.GetString("app_id", ""),
		"mch_id":           config.GetString("mch_id", ""),
		"partner_trade_no": partnerTradeNo,
	}

	endpoint := comp.Wrap("/mmpaymkttransfers/gettransferinfo")
	_, err := comp.SafeRequest(ctx, endpoint, params, http.MethodPost, &object.HashMap{}, nil, result)

	return result, err
}

// Send MerchantPay to balance.
// https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_2
func (comp *Client) ToBalance(ctx context.Context, data *request.RequestTransferToBalance) (*response.ResponseTransferToBalance, error) {

	result := &response.ResponseTransferToBalance{}

	config := (*comp.App).GetConfig()

	params, err := object.StructToHashMapWithXML(data)
	if err != nil {
		return nil, err
	}
	base := &object.HashMap{
		"mchid":     config.GetString("mch_id", ""),
		"mch_appid": config.GetString("app_id", ""),
	}

	params = object.MergeHashMap(params, base)
	endpoint := comp.Wrap("mmpaymkttransfers/promotion/transfers")
	_, err = comp.SafeRequest(ctx, endpoint, params, http.MethodPost, &object.HashMap{}, nil, result)

	return result, err
}

// Query MerchantPay order to BankCard.
// https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_3
func (comp *Client) QueryBankCardOrder(ctx context.Context, partnerTradeNo string) (*response.ResponseGetTransferInfo, error) {
	config := (*comp.App).GetConfig()

	result := &response.ResponseGetTransferInfo{}

	params := &object.HashMap{
		"mch_id":           config.GetString("mch_id", ""),
		"partner_trade_no": partnerTradeNo,
	}

	endpoint := comp.Wrap("/mmpaymkttransfers/query_bank")
	_, err := comp.SafeRequest(ctx, endpoint, params, http.MethodPost, &object.HashMap{}, nil, result)

	return result, err
}

// Send MerchantPay to BankCard.
// https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=24_2
func (comp *Client) ToBankCard(ctx context.Context, data *request.RequestToBankCard) (*response2.ResponsePayment, error) {

	result := &response2.ResponsePayment{}

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

	cipherData, err := rsaSigner.RSAEncryptor.Encrypt([]byte(data.EncBankNO))
	buffer := base64.StdEncoding.EncodeToString(cipherData)
	if err != nil {
		return nil, err
	}
	data.EncBankNO = string(buffer)
	if err != nil {
		return nil, err
	}

	cipherData, err = rsaSigner.RSAEncryptor.Encrypt([]byte(data.EncTrueName))
	buffer = base64.StdEncoding.EncodeToString(cipherData)
	if err != nil {
		return nil, err
	}
	data.EncTrueName = string(buffer)
	if err != nil {
		return nil, err
	}

	//params, err := object.StructToHashMapWithTag(data,"json")
	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	endpoint := comp.Wrap("mmpaysptrans/pay_bank")
	_, err = comp.SafeRequest(ctx, endpoint, params, http.MethodPost, &object.HashMap{}, nil, result)

	return result, err
}
