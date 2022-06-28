package code

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
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

// 上传小程序代码并生成体验版
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/commit.html
func (comp *Client) UpdateSignature(templateID string, extJson string, version string, description string) (*response2.ResponseOpenPlatform, error) {

	result := &response2.ResponseOpenPlatform{}

	params := &object.HashMap{
		"template_id":  templateID,
		"ext_json":     extJson,
		"user_version": version,
		"user_desc":    description,
	}
	_, err := comp.HttpPostJson("cgi-bin/wxa/commit", params, nil, nil, result)

	return result, err

}
