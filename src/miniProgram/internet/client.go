package internet

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type Client struct {
	*kernel.BaseClient
}

// 获取用户encryptKey。
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/internet/internet.getUserEncryptKey.html
func (comp *Client) GetUserEncryptKey(openID string, signature string, sigMethod string) (*response2.ResponseMiniProgram, error) {

	result := &response2.ResponseMiniProgram{}

	params := &object.StringMap{
		"openid":     openID,
		"signature":  signature,
		"sig_method": sigMethod,
	}

	_, err := comp.HttpPostJson("wxa/business/getuserencryptkey", nil, params, nil, result)

	return result, err
}
