package customer

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/work/accountService/customer/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

// 读取消息
// https://work.weixin.qq.com/api/doc/90000/90135/94670
func (comp *Client) BatchGet(externalUserIDList []string) (*response.ResponseCustomerBatchGet, error) {

	result := &response.ResponseCustomerBatchGet{}

	options := &object.HashMap{
		"external_userid_list": externalUserIDList,
	}

	_, err := comp.HttpPostJson("cgi-bin/kf/customer/batchget", options, nil, nil, result)

	return result, err
}



// 获取配置的专员与客户群
// https://work.weixin.qq.com/api/doc/90000/90135/94674
func (comp *Client) GetUpgradeServiceConfig() (*response.ResponseCustomerGetUpgradeServiceConfig, error) {

	result := &response.ResponseCustomerGetUpgradeServiceConfig{}


	_, err := comp.HttpGet("cgi-bin/kf/customer/get_upgrade_service_config", nil, nil, result)

	return result, err
}

