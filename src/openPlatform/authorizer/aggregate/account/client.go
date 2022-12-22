package account

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openPlatform/authorizer/aggregate/account/response"
)

type Client struct {
	*kernel.BaseClient
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

// 创建开放平台帐号并绑定公众号/小程序
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/account/create.html
func (comp *Client) Create() (*response.ResponseCreate, error) {

	result := &response.ResponseCreate{}

	config := (*comp.App).GetConfig()
	params := &object.HashMap{
		"appid": config.GetString("app_id", ""),
	}
	_, err := comp.HttpPostJson("cgi-bin/open/create", params, nil, nil, result)

	return result, err

}

// 将公众号/小程序绑定到开放平台帐号下
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/account/bind.html
func (comp *Client) BindTo(openAppID string) (*response2.ResponseOpenPlatform, error) {

	result := &response2.ResponseOpenPlatform{}

	config := (*comp.App).GetConfig()
	params := &object.HashMap{
		"appid":      config.GetString("app_id", ""),
		"open_appid": openAppID,
	}
	_, err := comp.HttpPostJson("cgi-bin/open/bind", params, nil, nil, result)

	return result, err

}

// 将公众号/小程序绑定到开放平台帐号下
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/account/unbind.html
func (comp *Client) UnbindFrom(openAppID string) (*response2.ResponseOpenPlatform, error) {

	result := &response2.ResponseOpenPlatform{}

	config := (*comp.App).GetConfig()
	params := &object.HashMap{
		"appid":      config.GetString("app_id", ""),
		"open_appid": openAppID,
	}
	_, err := comp.HttpPostJson("cgi-bin/open/unbind", params, nil, nil, result)

	return result, err

}

// 获取公众号/小程序所绑定的开放平台帐号
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/account/get.html
func (comp *Client) GetBinding() (*response.ResponseGetBinding, error) {

	result := &response.ResponseGetBinding{}

	config := (*comp.App).GetConfig()
	params := &object.HashMap{
		"appid": config.GetString("app_id", ""),
	}
	_, err := comp.HttpPostJson("cgi-bin/open/get", params, nil, nil, result)

	return result, err

}
