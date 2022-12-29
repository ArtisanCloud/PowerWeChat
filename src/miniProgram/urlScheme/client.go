package urlScheme

import (
	"context"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/urlScheme/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/urlScheme/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// 获取小程序 scheme 码，适用于短信、邮件、外部网页、微信内等拉起小程序的业务场景
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-scheme/urlscheme.generate.html
func (comp *Client) Generate(ctx context.Context, param *request.URLSchemeGenerate) (*response.ResponseUrlScheme, error) {

	result := &response.ResponseUrlScheme{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/generatescheme", param, nil, nil, result)

	return result, err
}
