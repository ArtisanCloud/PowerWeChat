package wxaCode

import (
	"github.com/ArtisanCloud/go-libs/http/response"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
)

type Client struct {
	*kernel.BaseClient
}

// 获取小程序二维码，适用于需要的码数量较少的业务场景
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.createQRCode.html
func (comp *Client) CreateQRCode(path string, width int) (*response.HttpResponse, error) {

	if width <= 0 {
		width = 430
	}

	data := &object.HashMap{
		"path":  path,
		"width": width,
	}

	rs, err := comp.HttpPostJson("cgi-bin/wxaapp/createwxaqrcode", data, nil, nil, nil)

	return rs.(*response.HttpResponse), err
}

// 获取小程序码，适用于需要的码数量较少的业务场景
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.get.html
func (comp *Client) Get(path string, width int,
	autoColor bool, lineColor *object.HashMap, isHyaline bool) (*response.HttpResponse, error) {

	if width <= 0 {
		width = 430
	}

	data := &object.HashMap{
		"path":       path,
		"width":      width,
		"auto_color": autoColor,
		"line_color": lineColor,
		"is_hyaline": isHyaline,
	}

	rs, err := comp.HttpPostJson("cgi-bin/wxa/getwxacode", data, nil, nil, nil)

	return rs.(*response.HttpResponse), err
}

// 获取小程序码，适用于需要的码数量极多的业务场景
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.getUnlimited.html
func (comp *Client) GetUnlimited(
	scene string, page string, width int,
	autoColor bool, lineColor *object.HashMap, isHyaline bool) (*response.HttpResponse, error) {

	if width <= 0 {
		width = 430
	}

	data := &object.HashMap{
		"scene":      scene,
		"page":       page,
		"width":      width,
		"auto_color": autoColor,
		"line_color": lineColor,
		"is_hyaline": isHyaline,
	}

	rs, err := comp.HttpPostJson("cgi-bin/wxa/getwxacodeunlimit", data, nil, nil, nil)

	return rs.(*response.HttpResponse), err
}
