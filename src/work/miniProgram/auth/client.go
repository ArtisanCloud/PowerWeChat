package auth

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/miniProgram/response"
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

// 第三方登录凭证校验
// https://developers.weixin.qq.com/miniprogram/dev/dev_wxwork/dev-doc/qywx-api/login/code2session.html
func (comp *Client) Session(ctx context.Context, code string) (*response.ResponseSession, error) {

	result := &response.ResponseSession{}

	_, err := comp.BaseClient.HttpGet(ctx, "cgi-bin/miniprogram/jscode2session", &object.StringMap{
		"js_code":    code,
		"grant_type": "authorization_code",
	}, nil, result)

	return result, err
}
