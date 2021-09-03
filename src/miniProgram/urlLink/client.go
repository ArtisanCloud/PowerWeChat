package urlScheme

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/miniProgram/urlLink/response"
)

type Client struct {
	*kernel.BaseClient
}

// 获取小程序 URL Link，适用于短信、邮件、网页、微信内等拉起小程序的业务场景
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-link/urllink.generate.html
func (comp *Client) Generate(
	path string, query string, isExpire bool,
	expireType int, expireInterval int, cloudBase *object.HashMap) (*response.ResponseURLLinkGenerate, error) {

	result := &response.ResponseURLLinkGenerate{}

	data := &object.HashMap{
		"path":            path,
		"query":           query,
		"is_expire":       isExpire,
		"expire_type":     expireType,
		"expire_interval": expireInterval,
		"cloud_base":      cloudBase,
	}

	_, err := comp.HttpPostJson("cgi-bin/wxa/generate_urllink", data, nil, nil, result)

	return result, err
}
