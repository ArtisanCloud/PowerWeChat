package urlLink

import (
	"github.com/ArtisanCloud/PowerLibs/v2/http/response"
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/miniProgram/urlLink/request"
	response4 "github.com/ArtisanCloud/PowerWeChat/v2/src/work/media/response"
	"net/http"
)

type Client struct {
	*kernel.BaseClient
}

// 获取小程序 URL Link，适用于短信、邮件、网页、微信内等拉起小程序的业务场景
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-link/urllink.generate.html
func (comp *Client) Generate(options *request.URLSchemeGenerate) (*http.Response, error) {

	var result string
	var header = &response4.ResponseHeaderMedia{}

	//data, err := object.StructToHashMapWithTag(options, "json")
	data, err := object.StructToHashMap(options)
	if err != nil {
		return nil, err
	}

	rs, err := comp.RequestRaw("cgi-bin/wxa/generate_urllink", "POST", data, &header, &result)

	httpRS := rs.(*response.HttpResponse).Response

	return httpRS, err
}
