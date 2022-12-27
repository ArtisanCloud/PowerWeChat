package externalPay

import (
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalPay/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalPay/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(&app, nil)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient,
	}, nil
}

// 新增商户号
// https://developer.work.weixin.qq.com/document/path/93666
func (comp *Client) AddMerchant(mchID string, merchantName string) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	options := &object.HashMap{
		"mch_id":        mchID,
		"merchant_name": merchantName,
	}

	_, err := comp.BaseClient.HttpPostJson("cgi-bin/externalpay/addmerchant", options, nil, nil, result)

	return result, err
}

// 查询商户号详情
// https://developer.work.weixin.qq.com/document/path/93666
func (comp *Client) GetMerchant(mchID string) (*response.ResponseExternalPayGetMerchant, error) {

	result := &response.ResponseExternalPayGetMerchant{}

	options := &object.HashMap{
		"mch_id": mchID,
	}

	_, err := comp.BaseClient.HttpPostJson("cgi-bin/externalpay/getmerchant", options, nil, nil, result)

	return result, err
}

// 删除商户号
// https://developer.work.weixin.qq.com/document/path/93666
func (comp *Client) DelMerchant(mchID string) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	options := &object.HashMap{
		"mch_id": mchID,
	}

	_, err := comp.BaseClient.HttpPostJson("cgi-bin/externalpay/delmerchant", options, nil, nil, result)

	return result, err
}

// 设置商户号使用范围
// https://developer.work.weixin.qq.com/document/path/93666
func (comp *Client) SetMchUseScope(options *request.RequestExternalPaySetMchUseScope) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.BaseClient.HttpPostJson("cgi-bin/externalpay/set_mch_use_scope", options, nil, nil, result)

	return result, err
}

// 获取对外收款记录
// https://developer.work.weixin.qq.com/document/path/93667
func (comp *Client) GetBillList(options *request.RequestExternalPayGetBillList) (*response.ResponseExternalPayGetBillList, error) {

	result := &response.ResponseExternalPayGetBillList{}

	_, err := comp.BaseClient.HttpPostJson("cgi-bin/externalpay/get_bill_list", options, nil, nil, result)

	return result, err
}
