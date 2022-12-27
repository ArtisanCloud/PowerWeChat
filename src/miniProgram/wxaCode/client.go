package wxaCode

import (
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	response4 "github.com/ArtisanCloud/PowerWeChat/v3/src/work/media/response"
	"io/fs"
	"io/ioutil"
	"net/http"
)

type Client struct {
	BaseClient *kernel.BaseClient
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

	rs, err := comp.BaseClient.RequestRaw("cgi-bin/wxaapp/createwxaqrcode", "POST", data, &header, &result)

	return rs, err

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

	rs, err := comp.BaseClient.RequestRaw("wxa/getwxacode", "POST", data, &header, &result)

	return rs, err
}

// 获取小程序码，适用于需要的码数量较少的业务场景
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.get.html
func (comp *Client) SaveAs(savedPath string, perm fs.FileMode,
	path string, width int64,
	autoColor bool, lineColor *power.HashMap, isHyaline bool) error {

	rs, err := comp.Get(path, width, autoColor, lineColor, isHyaline)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(rs.Body)
	err = ioutil.WriteFile(savedPath, body, perm) //保存图片地址

	return err
}

// 获取小程序码，适用于需要的码数量极多的业务场景
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.getUnlimited.html
func (comp *Client) GetUnlimited(
	scene string, page string,
	checkPath bool, envVersion string, width int64,
	autoColor bool, lineColor *power.HashMap, isHyaline bool) (*http.Response, error) {

	var result string
	var header = &response4.ResponseHeaderMedia{}

	if width <= 0 {
		width = 430
	}

	data := &object.HashMap{
		"form_params": &object.HashMap{
			"scene":       scene,
			"page":        page,
			"width":       width,
			"check_path":  checkPath,
			"env_version": envVersion,
			"auto_color":  autoColor,
			"line_color":  lineColor.ToHashMap(),
			"is_hyaline":  isHyaline,
		},
	}

	rs, err := comp.BaseClient.RequestRaw("wxa/getwxacodeunlimit", "POST", data, &header, &result)

	return rs, err
}
