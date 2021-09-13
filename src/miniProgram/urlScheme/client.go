package urlScheme

import (
	"github.com/ArtisanCloud/go-libs/http/response"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	response4 "github.com/ArtisanCloud/power-wechat/src/work/media/response"
	"net/http"
)

type Client struct {
	*kernel.BaseClient
}

// 获取小程序 scheme 码，适用于短信、邮件、外部网页、微信内等拉起小程序的业务场景
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-scheme/urlscheme.generate.html
func (comp *Client) Generate(
	jumpWxa *power.HashMap, isExpire bool, expireType int,
	expireTime int, expireInterval int) (*http.Response, error) {

	var result string
	var header = &response4.ResponseHeaderMedia{}

	data := &object.HashMap{
		"form_params": &object.HashMap{
			"jump_wxa":        jumpWxa.ToHashMap(),
			"is_expire":       isExpire,
			"expire_type":     expireType,
			"expire_time":     expireTime,
			"expire_interval": expireInterval,
		},
	}

	rs, err := comp.RequestRaw("wxa/generatescheme", "POST", data,  &header, &result)

	httpRS := rs.(*response.HttpResponse).Response

	return httpRS, err
}
