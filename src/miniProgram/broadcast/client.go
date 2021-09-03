package internet

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	response2 "github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type Client struct {
	*kernel.BaseClient
}

// 获取用户encryptKey。
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/internet/internet.getUserEncryptKey.html
func (comp *Client) GetUserEncryptKey(openID string, signature string, sigMethod string) (*response2.ResponseMiniProgram, error) {

	result := &response2.ResponseMiniProgram{}

	data := &object.HashMap{
		"openid":     openID,
		"signature":  signature,
		"sig_method": sigMethod,
	}

	_, err := comp.HttpPostJson("wxa/business/getuserencryptkey", data, nil, nil, result)

	return result, err
}
