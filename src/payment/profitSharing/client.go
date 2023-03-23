package profitSharing

import (
	"context"
	"errors"
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	payment "github.com/ArtisanCloud/PowerWeChat/v3/src/payment/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/profitSharing/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/profitSharing/response"
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

// Share Orders.
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_1.shtml
func (comp *Client) Share(ctx context.Context, param *request.RequestShare) (*response.ResponseProfitSharingOrder, error) {

	result := &response.ResponseProfitSharingOrder{}

	if param.AppID == "" {
		config := (*comp.App).GetConfig()
		param.AppID = config.GetString("app_id", "")
	}

	// 需要对接受者对名字加密
	if param.Receivers != nil && len(param.Receivers) > 0 {
		if comp.RsaOAEP == nil {
			return nil, errors.New("支付模块的RsaOAEP配置有误，要加密敏感信息，请确保正确配置EncryptOAEP的公钥路径")
		}

		for _, receiver := range param.Receivers {
			if receiver.Name != "" {
				buffer, err := comp.RsaOAEP.RSAEncryptor.Encrypt([]byte(receiver.Name))
				if err != nil {
					return nil, err
				}
				receiver.Name = string(buffer)
				if err != nil {
					return nil, err
				}
			}
		}

	}

	//options, err := object.StructToHashMapWithTag(param,"json")
	options, err := object.StructToHashMap(param)

	endpoint := comp.Wrap("/v3/profitsharing/orders")
	_, err = comp.Request(ctx, endpoint, nil, http.MethodPost, options, false, nil, result)

	return result, err
}

// Query Profit Sharing Result.
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_2.shtml
func (comp *Client) Query(ctx context.Context, transactionID string, outOrderNO string) (*response.ResponseProfitSharingOrder, error) {

	result := &response.ResponseProfitSharingOrder{}

	params := &object.StringMap{
		"transaction_id": transactionID,
	}

	endpoint := comp.Wrap(fmt.Sprintf("/v3/profitsharing/orders/%s", outOrderNO))
	_, err := comp.Request(ctx, endpoint, params, http.MethodPost, nil, false, nil, result)

	return result, err
}

// Share Return.
// https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_7&index=8
func (comp *Client) Return(ctx context.Context, data *request.RequestShareReturn) (*response.ResponseProfitSharingReturn, error) {

	result := &response.ResponseProfitSharingReturn{}

	config := (*comp.App).GetConfig()

	params, err := object.StructToHashMapWithXML(data)
	if err != nil {
		return nil, err
	}
	base := &object.HashMap{
		"return_amount": 1,
		"appid":         config.GetString("app_id", ""),
		"mch_id":        config.GetString("mch_id", ""),
	}
	params = object.MergeHashMap(params, base)

	endpoint := comp.Wrap("/secapi/pay/profitsharingreturn")
	_, err = comp.SafeRequest(ctx, endpoint, params, http.MethodPost, &object.HashMap{}, nil, result)

	return result, err
}

// Query Return Orders Result.
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_2.shtml

func (comp *Client) QueryReturn(ctx context.Context, outOrderNO string, outReturnNO string) (*response.ResponseProfitSharingReturnOrder, error) {

	result := &response.ResponseProfitSharingReturnOrder{}

	params := &object.StringMap{
		"out_order_no": outOrderNO,
	}

	endpoint := comp.Wrap(fmt.Sprintf("/v3/profitsharing/return-orders/%s", outReturnNO))
	_, err := comp.Request(ctx, endpoint, params, http.MethodPost, nil, false, nil, result)

	return result, err
}

// UnFreeze Orders.
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_5.shtml

func (comp *Client) UnfreezeOrders(ctx context.Context, transactionID string, outOrderNO string, description string) (*response.ResponseProfitSharingOrder, error) {

	result := &response.ResponseProfitSharingOrder{}

	options := &object.HashMap{
		"transaction_id": transactionID,
		"out_order_no":   outOrderNO,
		"description":    description,
	}

	endpoint := comp.Wrap("/v3/profitsharing/orders/unfreeze")
	_, err := comp.Request(ctx, endpoint, nil, http.MethodPost, options, false, nil, result)

	return result, err
}

// Query Transaction Amounts.
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_6.shtml

func (comp *Client) QueryTransactions(ctx context.Context, transactionID string) (*response.ResponseProfitSharingTransaction, error) {

	result := &response.ResponseProfitSharingTransaction{}

	endpoint := comp.Wrap(fmt.Sprintf("/v3/profitsharing/transactions/%s/amounts", transactionID))
	_, err := comp.Request(ctx, endpoint, nil, http.MethodPost, nil, false, nil, result)

	return result, err
}

// Add Receiver.
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_8.shtml

func (comp *Client) AddReceiver(
	ctx context.Context,
	receiverType string, account string, name string,
	relationType string, customRelation string) (*response.ResponseProfitSharingAddReceiver, error) {

	result := &response.ResponseProfitSharingAddReceiver{}

	if name != "" {
		if comp.RsaOAEP == nil {
			return nil, errors.New("支付模块的RsaOAEP配置有误，要加密敏感信息，请确保正确配置EncryptOAEP的公钥路径")
		}
		buffer, err := comp.RsaOAEP.RSAEncryptor.Encrypt([]byte(name))
		if err != nil {
			return nil, err
		}
		name = string(buffer)
		if err != nil {
			return nil, err
		}
	}

	options := &object.HashMap{
		"type":            receiverType,
		"account":         account,
		"name":            name,
		"relation_type":   relationType,
		"custom_relation": customRelation,
	}

	endpoint := comp.Wrap("/v3/profitsharing/receivers/add")
	_, err := comp.Request(ctx, endpoint, nil, http.MethodPost, options, false, nil, result)

	return result, err
}

// Delete Receiver.
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_9.shtml

func (comp *Client) DeleteReceiver(ctx context.Context, receiverType string, account string) (*response.ResponseProfitSharingDeleteReceiver, error) {

	result := &response.ResponseProfitSharingDeleteReceiver{}

	config := (*comp.App).GetConfig()
	options := &object.HashMap{
		"appid":   config.GetString("app_id", ""),
		"type":    receiverType,
		"account": account,
	}

	endpoint := comp.Wrap("/v3/profitsharing/receivers/delete")
	_, err := comp.Request(ctx, endpoint, nil, http.MethodPost, options, false, nil, result)

	return result, err
}

// Get Bills.
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_1_11.shtml

func (comp *Client) GetBills(ctx context.Context, subMchID string, billDate string, tarType string) (*response.ResponseProfitSharingGetBills, error) {

	result := &response.ResponseProfitSharingGetBills{}

	params := &object.StringMap{
		"sub_mchid": subMchID,
		"bill_date": billDate,
		"tar_type":  tarType,
	}

	endpoint := comp.Wrap("/v3/profitsharing/bills")
	_, err := comp.Request(ctx, endpoint, params, http.MethodPost, nil, false, nil, result)

	return result, err
}
