package setting

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/openPlatform/authorizer/miniProgram/setting/request"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/openPlatform/authorizer/miniProgram/setting/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app *kernel.ApplicationInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(app, nil)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient,
	}, nil
}

// 获取接口列表
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/apply_api/get_privacy_interface.html#请求地址
func (comp *Client) Get() (*response.ResponseGet, error) {

	result := &response.ResponseGet{}

	_, err := comp.HttpPostJson("wxa/security/get_privacy_interface", nil, nil, nil, result)

	return result, err

}

// 申请接口
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/apply_api/apply_privacy_interface.html
func (comp *Client) Set(params *request.RequestSet) (*response.ResponseSet, error) {

	result := &response.ResponseSet{}

	_, err := comp.HttpPostJson("wxa/security/apply_privacy_interface", params, nil, nil, result)

	return result, err

}
