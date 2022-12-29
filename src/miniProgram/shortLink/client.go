package shortLink

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/shortLink/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// 获取小程序 Short Link
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/short-link/shortlink.generate.html
func (comp *Client) Generate(ctx context.Context, pageUrl string, pageTitle string, isPermanent bool) (*response.ResponseShortLinkGenerate, error) {

	result := &response.ResponseShortLinkGenerate{}

	data := &object.HashMap{
		"page_url":     pageUrl,
		"page_title":   pageTitle,
		"is_permanent": isPermanent,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/genwxashortlink", data, nil, nil, result)

	return result, err
}
