package profitSharing

import (
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	payment "github.com/ArtisanCloud/PowerWeChat/v2/src/payment/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/payment/profitSharing/request"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/payment/profitSharing/response"
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
func (comp *Client) Share(param *request.RequestShare) (*response.ResponseProfitSharingOrder, error) {

	result := &response.ResponseProfitSharingOrder{}

	if param.AppID == "" {
		config := (*comp.App).GetConfig()
		param.AppID = config.GetString("app_id", "")
	}
	//options, err := object.StructToHashMapWithTag(param,"json")
	options, err := object.StructToHashMap(param)

	endpoint := comp.Wrap("/v3/profitsharing/orders")
	_, err = comp.Request(endpoint, nil, "POST", options, false, nil, result)

	return result, err
}

// Query Profit Sharing Result.
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_2.shtml
func (comp *Client) Query(transactionID string, outOrderNO string) (*response.ResponseProfitSharingOrder, error) {

	result := &response.ResponseProfitSharingOrder{}

	params := &object.StringMap{
		"transaction_id": transactionID,
	}

	endpoint := comp.Wrap(fmt.Sprintf("/v3/profitsharing/orders/%s", outOrderNO))
	_, err := comp.Request(endpoint, params, "GET", nil, false, nil, result)

	return result, err
}

// Query Return Orders Result.
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_2.shtml

func (comp *Client) QueryReturn(outOrderNO string, outReturnNO string) (*response.ResponseProfitSharingReturnOrder, error) {

	result := &response.ResponseProfitSharingReturnOrder{}

	params := &object.StringMap{
		"out_order_no": outOrderNO,
	}

	endpoint := comp.Wrap(fmt.Sprintf("/v3/profitsharing/return-orders/%s", outReturnNO))
	_, err := comp.Request(endpoint, params, "GET", nil, false, nil, result)

	return result, err
}

// UnFreeze Orders.
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_5.shtml

func (comp *Client) UnfreezeOrders(transactionID string, outOrderNO string, description string) (*response.ResponseProfitSharingOrder, error) {

	result := &response.ResponseProfitSharingOrder{}

	options := &object.HashMap{
		"transaction_id": transactionID,
		"out_order_no":   outOrderNO,
		"description":    description,
	}

	endpoint := comp.Wrap("/v3/profitsharing/orders/unfreeze")
	_, err := comp.Request(endpoint, nil, "POST", options, false, nil, result)

	return result, err
}

// Query Transaction Amounts.
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_6.shtml

func (comp *Client) QueryTransactions(transactionID string) (*response.ResponseProfitSharingTransaction, error) {

	result := &response.ResponseProfitSharingTransaction{}

	endpoint := comp.Wrap(fmt.Sprintf("/v3/profitsharing/transactions/%s/amounts", transactionID))
	_, err := comp.Request(endpoint, nil, "GET", nil, false, nil, result)

	return result, err
}

// Add Receiver.
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_8.shtml

func (comp *Client) AddReceiver(
	receiverType string, account string, name string,
	relationType string, customRelation string) (*response.ResponseProfitSharingAddReceiver, error) {

	result := &response.ResponseProfitSharingAddReceiver{}

	options := &object.HashMap{
		"type":            receiverType,
		"accountService":  account,
		"name":            name,
		"relation_type":   relationType,
		"custom_relation": customRelation,
	}

	endpoint := comp.Wrap("/v3/profitsharing/receivers/add")
	_, err := comp.Request(endpoint, nil, "POST", options, false, nil, result)

	return result, err
}

// Delete Receiver.
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_9.shtml

func (comp *Client) DeleteReceiver(receiverType string, account string) (*response.ResponseProfitSharingDeleteReceiver, error) {

	result := &response.ResponseProfitSharingDeleteReceiver{}

	config := (*comp.App).GetConfig()
	options := &object.HashMap{
		"appid":          config.GetString("app_id", ""),
		"type":           receiverType,
		"accountService": account,
	}

	endpoint := comp.Wrap("/v3/profitsharing/receivers/delete")
	_, err := comp.Request(endpoint, nil, "POST", options, false, nil, result)

	return result, err
}

// Get Bills.
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_1_11.shtml

func (comp *Client) GetBills(subMchID string, billDate string, tarType string) (*response.ResponseProfitSharingGetBills, error) {

	result := &response.ResponseProfitSharingGetBills{}

	params := &object.StringMap{
		"sub_mchid": subMchID,
		"bill_date": billDate,
		"tar_type":  tarType,
	}

	endpoint := comp.Wrap("/v3/profitsharing/bills")
	_, err := comp.Request(endpoint, params, "GET", nil, false, nil, result)

	return result, err
}
