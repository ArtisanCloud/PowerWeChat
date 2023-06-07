package wxaCode

import (
	"context"
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
func (comp *Client) CreateQRCode(ctx context.Context, path string, width int64) (*http.Response, error) {

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

	rs, err := comp.BaseClient.RequestRaw(ctx, "cgi-bin/wxaapp/createwxaqrcode", http.MethodPost, data, &header, nil)

	return rs, err

}

// 获取小程序码，适用于需要的码数量较少的业务场景
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.get.html
func (comp *Client) Get(ctx context.Context, path string, width int64,
	autoColor bool, lineColor *power.HashMap, isHyaline bool) (*http.Response, error) {

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

	rs, err := comp.BaseClient.RequestRaw(ctx, "wxa/getwxacode", http.MethodPost, data, &header, nil)

	return rs, err
}

// 获取小程序码，适用于需要的码数量较少的业务场景
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.get.html
func (comp *Client) SaveAs(ctx context.Context, savedPath string, perm fs.FileMode,
	path string, width int64,
	autoColor bool, lineColor *power.HashMap, isHyaline bool) error {

	rs, err := comp.Get(ctx, path, width, autoColor, lineColor, isHyaline)
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
	ctx context.Context,
	scene string, page string,
	checkPath bool, envVersion string, width int64,
	autoColor bool, lineColor *power.HashMap, isHyaline bool) (*http.Response, error) {

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

	rs, err := comp.BaseClient.RequestRaw(ctx, "wxa/getwxacodeunlimit", http.MethodPost, data, &header, nil)

	return rs, err
}
