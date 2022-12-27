package privacy

import (
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openPlatform/authorizer/miniProgram/privacy/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openPlatform/authorizer/miniProgram/privacy/response"
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

// 查询小程序用户隐私保护指引
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/privacy_config/get_privacy_setting.html#请求地址
func (comp *Client) Get() (*response.ResponseGet, error) {

	result := &response.ResponseGet{}

	_, err := comp.BaseClient.HttpPostJson("cgi-bin/component/getprivacysetting", nil, nil, nil, result)

	return result, err

}

// 配置小程序用户隐私保护指引
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/privacy_config/set_privacy_setting.html
func (comp *Client) Set(params *request.RequestSet) (*response2.ResponseOpenPlatform, error) {

	result := &response2.ResponseOpenPlatform{}

	_, err := comp.BaseClient.HttpPostJson("cgi-bin/component/setprivacysetting", params, nil, nil, result)

	return result, err

}

// 上传小程序用户隐私保护指引
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/privacy_config/upload_privacy_exfile.html#请求地址
func (comp *Client) Upload(path string) (*response.ResponseUpload, error) {

	result := &response.ResponseUpload{}

	var files *object.HashMap
	if path != "" {
		files = &object.HashMap{
			"file": path,
		}
	}

	_, err := comp.BaseClient.HttpUpload("cgi-bin/component/uploadprivacyextfile", files, nil, nil, nil, result)

	return result, err

}
