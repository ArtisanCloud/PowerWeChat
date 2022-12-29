package security

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/security/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// 向插件开发者发起使用插件的申请
// https://developers.weixin.qq.com/miniprogram/dev/framework/security.imgSecCheck.html
func (comp *Client) ImgSecCheck(ctx context.Context, path string, form *power.HashMap) (*response2.ResponseMiniProgram, error) {

	result := &response2.ResponseMiniProgram{}

	var files *object.HashMap
	if path != "" {
		files = &object.HashMap{
			"media": path,
		}
	}

	var formData *kernel.UploadForm
	if form != nil {
		formData = &kernel.UploadForm{
			FileName: (*form)["name"].(string),
			Contents: []*kernel.UploadContent{
				&kernel.UploadContent{
					Name:  "media",
					Value: (*form)["value"],
				},
			},
		}
	}

	_, err := comp.BaseClient.HttpUpload(ctx, "wxa/img_sec_check", files, formData, nil, nil, result)

	return result, err
}

// 异步校验图片/音频是否含有违法违规内容
// https://developers.weixin.qq.com/miniprogram/dev/framework/security.mediaCheckAsync-v1.html#请求地址
func (comp *Client) MediaCheckAsync(ctx context.Context, mediaURL string, mediaType int, version int, openID string, scene int) (*response.ResponseSecurityMediaCheckASync, error) {

	result := &response.ResponseSecurityMediaCheckASync{}

	data := &object.HashMap{
		"media_url":  mediaURL,
		"media_type": mediaType,
		"version":    version,
		"openid":     openID,
		"scene":      scene,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/media_check_async", data, nil, nil, result)

	return result, err
}

// 检查一段文本是否含有违法违规内容
// https://developers.weixin.qq.com/miniprogram/dev/framework/security.msgSecCheck-v1.html#HTTPS%20调用
func (comp *Client) MsgSecCheck(
	ctx context.Context,
	openID string, scene int, version int, content string,
	nickname string, title string, signature string) (*response.ResponseSecurityMsgCheckASync, error) {

	result := &response.ResponseSecurityMsgCheckASync{}

	data := &object.HashMap{
		"openid":    openID,
		"scene":     scene,
		"version":   version,
		"content":   content,
		"nickname":  nickname,
		"title":     title,
		"signature": signature,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/msg_sec_check", data, nil, nil, result)

	return result, err
}
