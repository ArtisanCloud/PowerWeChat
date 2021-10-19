package urlLink

import (
	"github.com/ArtisanCloud/PowerLibs/http/response"
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	response4 "github.com/ArtisanCloud/PowerWeChat/src/work/media/response"
	"net/http"
)

type Client struct {
	*kernel.BaseClient
}

// 获取小程序 URL Link，适用于短信、邮件、网页、微信内等拉起小程序的业务场景
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-link/urllink.generate.html
func (comp *Client) Generate(
	path string, query string, isExpire bool,
	expireType int, expireInterval int, cloudBase *power.HashMap) (*http.Response, error) {

	var result string
	var header = &response4.ResponseHeaderMedia{}

	data := &object.HashMap{
		"path":            path,
		"query":           query,
		"is_expire":       isExpire,
		"expire_type":     expireType,
		"expire_interval": expireInterval,
		"cloud_base":      cloudBase.ToHashMap(),
	}

	rs, err := comp.RequestRaw("cgi-bin/wxa/generate_urllink", "POST", data, &header, &result)

	httpRS := rs.(*response.HttpResponse).Response

	return httpRS, err
}
