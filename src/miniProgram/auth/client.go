package auth

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/powerwechat/src/kernel"
	"github.com/ArtisanCloud/powerwechat/src/miniProgram/auth/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app *kernel.ApplicationInterface) *Client {

	return &Client{
		kernel.NewBaseClient(app, nil),
	}
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
