package urlLink

import (
	"github.com/ArtisanCloud/go-libs/http/response"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
)

type Client struct {
	*kernel.BaseClient
}

// 获取小程序 scheme 码，适用于短信、邮件、外部网页、微信内等拉起小程序的业务场景
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-scheme/urlscheme.generate.html
func (comp *Client) Generate(
	jumpWxa *object.HashMap, isExpire bool, expireType int,
	expireTime int, expireInterval int) (*response.HttpResponse, error) {

	data := &object.HashMap{
		"jump_wxa":        jumpWxa,
		"is_expire":       isExpire,
		"expire_type":     expireType,
		"expire_time":     expireTime,
		"expire_interval": expireInterval,
	}

	rs, err := comp.HttpPostJson("cgi-bin/wxa/generatescheme", data, nil, nil, nil)

	return rs.(*response.HttpResponse), err
}
