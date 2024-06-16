package auth

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/support"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/auth/response"
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

// 登录凭证校验。
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html
func (comp *Client) Session(ctx context.Context, code string) (*response.ResponseCode2Session, error) {

	result := &response.ResponseCode2Session{}

	config := (*comp.BaseClient.App).GetConfig()

	params := &object.StringMap{
		"appid":      config.GetString("app_id", ""),
		"secret":     config.GetString("secret", ""),
		"js_code":    code,
		"grant_type": "authorization_code",
	}

	_, err := comp.BaseClient.HttpGet(ctx, "sns/jscode2session", params, nil, result)

	return result, err
}

// 检验登录态。
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-login/checkSessionKey.html
func (comp *Client) CheckSession(ctx context.Context, openId string, sessionKey string) (*response.ResponseCode2Session, error) {

	result := &response.ResponseCode2Session{}

	config := (*comp.BaseClient.App).GetConfig()

	sign, err := support.SignSHA256WithHMac([]byte(sessionKey), "")
	if err != nil {
		return nil, err
	}
	params := &object.StringMap{
		"appid":      config.GetString("app_id", ""),
		"secret":     config.GetString("secret", ""),
		"openid":     openId,
		"signature":  string(sign),
		"sig_method": "hmac_sha256",
	}

	_, err = comp.BaseClient.HttpGet(ctx, "wxa/checksession", params, nil, result)

	return result, err
}

// 重置登录态。
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-login/ResetUserSessionKey.html
func (comp *Client) ResetUserSessionKey(ctx context.Context, openId string, sessionKey string) (*response2.ResponseMiniProgram, error) {

	result := &response2.ResponseMiniProgram{}

	config := (*comp.BaseClient.App).GetConfig()

	sign, err := support.SignSHA256WithHMac([]byte(sessionKey), "")
	if err != nil {
		return nil, err
	}
	params := &object.StringMap{
		"appid":      config.GetString("app_id", ""),
		"secret":     config.GetString("secret", ""),
		"openid":     openId,
		"signature":  string(sign),
		"sig_method": "hmac_sha256",
	}

	_, err = comp.BaseClient.HttpGet(ctx, "wxa/resetusersessionkey", params, nil, result)

	return result, err
}
