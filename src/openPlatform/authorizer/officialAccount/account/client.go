package account

import (
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openPlatform/authorizer/aggregate/account"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openPlatform/authorizer/officialAccount/account/response"
)

type Client struct {
	*account.Client

	// PowerWechat\OpenPlatform\Application
	component kernel.ApplicationInterface
}

func NewClient(app kernel.ApplicationInterface, component kernel.ApplicationInterface) (*Client, error) {
	client, err := account.NewClient(app)
	if err != nil {
		return nil, err
	}
	return &Client{
		client,
		component,
	}, nil
}

// 从第三方平台跳转至微信公众平台授权注册页面, 授权注册小程序.
func (comp *Client) GetFastRegistrationUrl(callbackUrl string, copyWxVerify bool) string {

	config := (*comp.App).GetConfig()
	componentConfig := comp.component.GetConfig()

	queries := &object.StringMap{
		"copy_wx_verify":  fmt.Sprintf("%t", copyWxVerify),
		"component_appid": componentConfig.GetString("app_id", ""),
		"appid":           config.GetString("app_id", ""),
		"redirect_uri":    callbackUrl,
	}

	return "https://mp.weixin.qq.com/cgi-bin/fastregisterauth?" + object.GetJoinedWithKSort(queries)

}

// 复用公众号主体快速注册小程序
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Register_Mini_Programs/fast_registration_of_mini_program.html
func (comp *Client) Register(ticket string) (*response.ResponseRegister, error) {

	result := &response.ResponseRegister{}

	_, err := comp.HttpPostJson("cgi-bin/account/fastregister", &object.HashMap{
		"ticket": ticket,
	}, nil, nil, result)

	return result, err

}
