package tester

import (
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openPlatform/authorizer/miniProgram/tester/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openPlatform/authorizer/miniProgram/tester/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
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

// 绑定微信用户为体验者
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_AdminManagement/Admin.html
func (comp *Client) Bind(wechatID string) (*response.ResponseBind, error) {

	result := &response.ResponseBind{}

	_, err := comp.BaseClient.HttpPostJson("wxa/bind_tester", &object.HashMap{
		"wechatid": wechatID,
	}, nil, nil, result)

	return result, err

}

// 解除绑定体验者
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_AdminManagement/unbind_tester.html
func (comp *Client) Unbind(params *request.RequestUnbind) (*response2.ResponseOpenPlatform, error) {

	result := &response2.ResponseOpenPlatform{}

	_, err := comp.BaseClient.HttpPostJson("wxa/unbind_tester", params, nil, nil, result)

	return result, err

}

// 获取体验者列表
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_AdminManagement/memberauth.html
func (comp *Client) List() (*response.ResponseList, error) {

	result := &response.ResponseList{}

	_, err := comp.BaseClient.HttpPostJson("wxa/memberauth", &object.HashMap{
		"action": "get_experiencer",
	}, nil, nil, result)

	return result, err
}
