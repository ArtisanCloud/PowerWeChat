package image

import (
	"context"
	"errors"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/image/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// 本接口提供基于小程序的图片智能裁剪能力
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/img/img.aiCrop.html
func (comp *Client) AICrop(ctx context.Context, imgURL string, img *power.HashMap) (*response.ResponseIMGAICrop, error) {

	result := &response.ResponseIMGAICrop{}

	_, err := comp.UploadImage(ctx, "cv/img/aicrop", imgURL, img, result)

	return result, err

}

// 本接口提供基于小程序的条码/二维码识别的API
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/img/img.scanQRCode.html
func (comp *Client) ScanQRCode(ctx context.Context, imgURL string, img *power.HashMap) (*response.ResponseIMGScanQRCode, error) {

	result := &response.ResponseIMGScanQRCode{}

	_, err := comp.UploadImage(ctx, "cv/img/aicrop", imgURL, img, result)

	return result, err

}

// 本接口提供基于小程序的图片高清化能力
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/img/img.superresolution.html
func (comp *Client) SuperResolution(ctx context.Context, imgURL string, img *power.HashMap) (*response.ResponseIMGSuperResolution, error) {

	result := &response.ResponseIMGSuperResolution{}

	_, err := comp.UploadImage(ctx, "cv/img/superresolution", imgURL, img, result)

	return result, err

}

func (comp *Client) UploadImage(ctx context.Context, entryPoint string, imgURL string, img *power.HashMap, result interface{}) (interface{}, error) {

	params := &object.StringMap{}
	var formData *kernel.UploadForm

	if imgURL != "" {
		params = &object.StringMap{
			"img_url": imgURL,
		}
	} else if img != nil {
		formData = &kernel.UploadForm{
			Contents: []*kernel.UploadContent{
				&kernel.UploadContent{
					Name:  (*img)["name"].(string),
					Value: (*img)["value"],
				},
			},
		}
	} else {
		return nil, errors.New("Please send image url or image form data ")
	}

	rs, err := comp.BaseClient.HttpUpload(ctx, entryPoint, nil, formData, params, nil, result)

	return rs, err
}
