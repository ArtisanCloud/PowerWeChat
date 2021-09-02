package transfer

import (
	"errors"
	"fmt"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	payment "github.com/ArtisanCloud/power-wechat/src/payment/kernel"
	"github.com/ArtisanCloud/power-wechat/src/payment/transfer/response"
	"net/http"
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
func (comp *Client) QueryBalanceOrder(partnerTradeNo string) (interface{}, error) {
	config := (*comp.App).GetConfig()

	options := &object.HashMap{
		"appid":            config.GetString("app_id", ""),
		"mch_id":           config.GetString("mch_id", ""),
		"partner_trade_no": partnerTradeNo,
	}

	endpoint := comp.Wrap("/mmpaymkttransfers/gettransferinfo")
	result, err := comp.SafeRequest(endpoint, nil, "POST", options, false, nil)

	return result, err
}

// Send MerchantPay to balance.
// https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_2

func (comp *Client) ToBalance(params *power.HashMap) (interface{}, error) {

	result := &response.ResponseTransfer{}

	config := (*comp.App).GetConfig()

	externalRequest := (*comp.App).GetComponent("ExternalRequest").(*http.Request)
	if (*params)["spbill_create_ip"] == nil || (*params)["spbill_create_ip"].(string) != "" {
		(*params)["spbill_create_ip"] = externalRequest.Host
	}

	base := &object.HashMap{
		"mch_id":    nil,
		"mchid":     config.GetString("mch_id", ""),
		"mch_appid": config.GetString("app_id", ""),
	}

	options := object.MergeHashMap(base, params.ToHashMap())

	endpoint := comp.Wrap("mmpaymkttransfers/promotion/transfers")
	_, err := comp.Request(endpoint, nil, "POST", options, false, nil, result)

	return result, err
}


// Query MerchantPay order to BankCard.
// https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_3
func (comp *Client) QueryBankCardOrder(partnerTradeNo string) (interface{}, error) {
	config := (*comp.App).GetConfig()

	options := &object.HashMap{
		"mch_id":           config.GetString("mch_id", ""),
		"partner_trade_no": partnerTradeNo,
	}

	endpoint := comp.Wrap("/mmpaymkttransfers/query_bank")
	result, err := comp.SafeRequest(endpoint, nil, "POST", options, false, nil)

	return result, err
}

// Send MerchantPay to BankCard.
// https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_2
func (comp *Client) ToBankCard(params *object.HashMap) (interface{}, error) {

	result := &response.ResponseTransfer{}

	//config := (*comp.App).GetConfig()

	for _, key := range []interface{}{
		(*params)["bank_code"],
		(*params)["partner_trade_no"],
		(*params)["enc_bank_no"],
		(*params)["enc_true_name"],
		(*params)["amount"],
	} {
		if key == nil {
			return nil, errors.New(fmt.Sprintf("\"%s\" is required.", key))
		}
	}

	//publicKey, err := ioutil.ReadFile(config.GetString("rsa_public_key_path","")
	//if err != nil {
	//	return nil, err
	//}

	//(*params)["enc_bank_no"] = RSAPublicEncrypt((*params)["enc_bank_no"].(string), publicKey),
	//(*params)["enc_true_name"] = RSAPublicEncrypt((*params)["enc_true_name"].(string), publicKey),

	endpoint := comp.Wrap("mmpaysptrans/pay_bank")
	_, err := comp.Request(endpoint, nil, "POST", params, false, nil, result)

	return result, err
}
