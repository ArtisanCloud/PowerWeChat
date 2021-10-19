package externalPay

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/powerwechat/src/kernel"
	response2 "github.com/ArtisanCloud/powerwechat/src/kernel/response"
	"github.com/ArtisanCloud/powerwechat/src/work/externalPay/request"
	"github.com/ArtisanCloud/powerwechat/src/work/externalPay/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

// 新增商户号
// https://work.weixin.qq.com/api/doc/90000/90135/93666
func (comp *Client) AddMerchant(mchID string, merchantName string) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	options := &object.HashMap{
		"mch_id":        mchID,
		"merchant_name": merchantName,
	}

	_, err := comp.HttpPostJson("cgi-bin/externalpay/addmerchant", options, nil, nil, result)

	return result, err
}

// 查询商户号详情
// https://work.weixin.qq.com/api/doc/90000/90135/93666
func (comp *Client) GetMerchant(mchID string) (*response.ResponseExternalPayGetMerchant, error) {

	result := &response.ResponseExternalPayGetMerchant{}

	options := &object.HashMap{
		"mch_id": mchID,
	}

	_, err := comp.HttpPostJson("cgi-bin/externalpay/getmerchant", options, nil, nil, result)

	return result, err
}

// 删除商户号
// https://work.weixin.qq.com/api/doc/90000/90135/93666
func (comp *Client) DelMerchant(mchID string) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	options := &object.HashMap{
		"mch_id": mchID,
	}

	_, err := comp.HttpPostJson("cgi-bin/externalpay/delmerchant", options, nil, nil, result)

	return result, err
}

// 设置商户号使用范围
// https://work.weixin.qq.com/api/doc/90000/90135/93666
func (comp *Client) SetMchUseScope(options *request.RequestExternalPaySetMchUseScope) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/externalpay/set_mch_use_scope", options, nil, nil, result)

	return result, err
}

// 获取对外收款记录
// https://work.weixin.qq.com/api/doc/90000/90135/93667
func (comp *Client) GetBillList(options *request.RequestExternalPayGetBillList) (*response.ResponseExternalPayGetBillList, error) {

	result := &response.ResponseExternalPayGetBillList{}

	_, err := comp.HttpPostJson("cgi-bin/externalpay/get_bill_list", options, nil, nil, result)

	return result, err
}
