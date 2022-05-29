package urlScheme

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	response2 "github.com/ArtisanCloud/PowerWeChat/v2/src/miniProgram/urlScheme/response"
)

type Client struct {
	*kernel.BaseClient
}

// 获取小程序 scheme 码，适用于短信、邮件、外部网页、微信内等拉起小程序的业务场景
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-scheme/urlscheme.generate.html
func (comp *Client) Generate(
	jumpWxa *power.HashMap, isExpire bool, expireType int,
	expireTime int, expireInterval int) (*response2.ResponseUrlScheme, error) {

	result := &response2.ResponseUrlScheme{}

	data := &object.HashMap{
		"jump_wxa":        jumpWxa.ToHashMap(),
		"is_expire":       isExpire,
		"expire_type":     expireType,
		"expire_time":     expireTime,
		"expire_interval": expireInterval,
	}

	_, err := comp.HttpPostJson("wxa/generatescheme", data, nil, nil, result)

	return result, err
}
