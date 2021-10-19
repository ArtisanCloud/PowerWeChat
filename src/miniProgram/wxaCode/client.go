package wxaCode

import (
	"github.com/ArtisanCloud/go-libs/http/response"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/powerwechat/src/kernel"
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	response4 "github.com/ArtisanCloud/powerwechat/src/work/media/response"
	"net/http"
)

type Client struct {
	*kernel.BaseClient
}

// 获取小程序二维码，适用于需要的码数量较少的业务场景
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.createQRCode.html
func (comp *Client) CreateQRCode(path string, width int64) (*http.Response, error) {

	var result string
	var header = &response4.ResponseHeaderMedia{}

	if width <= 0 {
		width = 430
	}

	data := &object.HashMap{
		"form_params": &object.HashMap{
			"path":  path,
			"width": width,
		},
	}

	rs, err := comp.RequestRaw("cgi-bin/wxaapp/createwxaqrcode", "POST", data, &header, &result)

	httpRS := rs.(*response.HttpResponse).Response

	return httpRS, err

}

// 获取小程序码，适用于需要的码数量较少的业务场景
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.get.html
func (comp *Client) Get(path string, width int64,
	autoColor bool, lineColor *power.HashMap, isHyaline bool) (*http.Response, error) {

	var result string
	var header = &response4.ResponseHeaderMedia{}

	if width <= 0 {
		width = 430
	}

	data := &object.HashMap{
		"form_params": &object.HashMap{
			"path":       path,
			"width":      width,
			"auto_color": autoColor,
			"line_color": lineColor.ToHashMap(),
			"is_hyaline": isHyaline,
		},
	}

	rs, err := comp.RequestRaw("wxa/getwxacode", "POST", data,  &header, &result)

	httpRS := rs.(*response.HttpResponse).Response

	return httpRS, err
}

// 获取小程序码，适用于需要的码数量极多的业务场景
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.getUnlimited.html
func (comp *Client) GetUnlimited(
	scene string, page string, width int64,
	autoColor bool, lineColor *power.HashMap, isHyaline bool) (*http.Response, error) {

	var result string
	var header = &response4.ResponseHeaderMedia{}

	if width <= 0 {
		width = 430
	}

	data := &object.HashMap{
		"form_params": &object.HashMap{
			"scene":      scene,
			"page":       page,
			"width":      width,
			"auto_color": autoColor,
			"line_color": lineColor.ToHashMap(),
			"is_hyaline": isHyaline,
		},
	}

	rs, err := comp.RequestRaw("wxa/getwxacodeunlimit", "POST",data, &header, &result)

	httpRS := rs.(*response.HttpResponse).Response

	return httpRS, err
}
