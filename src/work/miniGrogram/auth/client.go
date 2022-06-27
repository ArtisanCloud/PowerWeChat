package auth

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/work/miniGrogram/response"
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

// 第三方登录凭证校验
// https://developers.weixin.qq.com/miniprogram/dev/dev_wxwork/dev-doc/qywx-api/login/code2session.html
func (comp *Client) Session(code string) (*response.ResponseSession, error) {

	result := &response.ResponseSession{}

	_, err := comp.HttpGet("cgi-bin/miniprogram/jscode2session", &object.StringMap{
		"js_code":    code,
		"grant_type": "authorization_code",
	}, nil, result)

	return result, err
}
