package shortLink

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/miniProgram/shortLink/response"
)

type Client struct {
	*kernel.BaseClient
}

// 获取小程序 Short Link
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/short-link/shortlink.generate.html
func (comp *Client) Generate(pageUrl string, pageTitle string, isPermanent bool) (*response.ResponseShortLinkGenerate, error) {

	result := &response.ResponseShortLinkGenerate{}

	data := &object.HashMap{
		"page_url":     pageUrl,
		"page_title":   pageTitle,
		"is_permanent": isPermanent,
	}

	_, err := comp.HttpPostJson("wxa/genwxashortlink", data, nil, nil, result)

	return result, err
}
