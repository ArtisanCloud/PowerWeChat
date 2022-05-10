package transfer

import (
	"crypto/sha256"
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/support"
	payment "github.com/ArtisanCloud/PowerWeChat/src/payment/kernel"
	"github.com/ArtisanCloud/PowerWeChat/src/payment/transfer/request"
	"github.com/ArtisanCloud/PowerWeChat/src/payment/transfer/response"
)

type Client struct {
	*payment.BaseClient
}

func NewClient(app *payment.ApplicationPaymentInterface) *Client {
	return &Client{
		payment.NewBaseClient(app),
	}
}

// Query MerchantPay to balance.
// https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_3
func (comp *Client) QueryBalanceOrder(partnerTradeNo string) (*response.ResponseGetTransferInfo, error) {
	config := (*comp.App).GetConfig()

	result := &response.ResponseGetTransferInfo{}

	options := &object.HashMap{
		"appid":            config.GetString("app_id", ""),
		"mch_id":           config.GetString("mch_id", ""),
		"partner_trade_no": partnerTradeNo,
	}

	endpoint := comp.Wrap("/mmpaymkttransfers/gettransferinfo")
	_, err := comp.SafeRequest(endpoint, nil, "POST", options, false, result)

	return result, err
}

// Send MerchantPay to balance.
// https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_2
func (comp *Client) ToBalance(params *request.RequestTransferToBalance) (*response.ResponseTransferToBalance, error) {

	result := &response.ResponseTransferToBalance{}

	config := (*comp.App).GetConfig()

	externalRequest := (*comp.App).GetExternalRequest()
	if params.SpBillCreateIP == "" {
		params.SpBillCreateIP = externalRequest.Host
	}

	options, err := object.StructToStringMap(params)
	base := &object.StringMap{
		//"mch_id":    "",
		"mchid":     config.GetString("mch_id", ""),
		"mch_appid": config.GetString("app_id", ""),
	}

	options = object.MergeStringMap(base, options)
	endpoint := comp.Wrap("mmpaymkttransfers/promotion/transfers")
	_, err = comp.SafeRequest(endpoint, nil, "POST", options, nil, result)

	return result, err
}

// Query MerchantPay order to BankCard.
// https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_3
func (comp *Client) QueryBankCardOrder(partnerTradeNo string) (*response.ResponseGetTransferInfo, error) {
	config := (*comp.App).GetConfig()

	result := &response.ResponseGetTransferInfo{}

	options := &object.HashMap{
		"mch_id":           config.GetString("mch_id", ""),
		"partner_trade_no": partnerTradeNo,
	}

	endpoint := comp.Wrap("/mmpaymkttransfers/query_bank")
	_, err := comp.SafeRequest(endpoint, nil, "POST", options, false, result)

	return result, err
}

// Send MerchantPay to BankCard.
// https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=24_2
func (comp *Client) ToBankCard(params *request.RequestToBankCard) (*response.ResponseTransfer, error) {

	result := &response.ResponseTransfer{}

	config := (*comp.App).GetConfig()

	rsaHandle := &support.RSAOaep{
		PublicKeyPath:  config.GetString("rsa_public_key_path", ""),
		PrivateKeyPath: config.GetString("rsa_private_key_path", ""),
	}

	bufferTemp, err := rsaHandle.DecryptOAEP(sha256.New(), []byte(params.EncBankNO))
	if err != nil {
		return nil, err
	}
	params.EncBankNO = string(bufferTemp)
	bufferTemp, err = rsaHandle.DecryptOAEP(sha256.New(), []byte(params.EncTrueName))
	if err != nil {
		return nil, err
	}
	params.EncTrueName = string(bufferTemp)

	endpoint := comp.Wrap("mmpaysptrans/pay_bank")
	_, err = comp.SafeRequest(endpoint, nil, "POST", params, nil, result)

	return result, err
}
