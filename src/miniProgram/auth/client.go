package auth

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/auth/response"
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

// 登录凭证校验。
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html
func (comp *Client) Session(code string) (*response.ResponseCode2Session, error) {

	result := &response.ResponseCode2Session{}

	config := (*comp.App).GetConfig()

	params := &object.StringMap{
		"appid":      config.GetString("app_id", ""),
		"secret":     config.GetString("secret", ""),
		"js_code":    code,
		"grant_type": "authorization_code",
	}

	_, err := comp.HttpGet("sns/jscode2session", params, nil, result)

	return result, err
}
