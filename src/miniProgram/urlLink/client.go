package urlLink

import (
	"context"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/urlLink/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/urlLink/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// 获取小程序 URL Link，适用于短信、邮件、网页、微信内等拉起小程序的业务场景
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-link/urllink.generate.html
func (comp *Client) Generate(ctx context.Context, options *request.URLLinkGenerate) (*response.ResponseURLLinkGenerate, error) {

	result := &response.ResponseURLLinkGenerate{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/generate_urllink", options, nil, nil, &result)

	return result, err
}
