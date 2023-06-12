package merchant

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	payment "github.com/ArtisanCloud/PowerWeChat/v3/src/payment/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/merchant/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/merchant/response"
)

type Client struct {
	*payment.BaseClient
}

func NewClient(app *payment.ApplicationPaymentInterface) (*Client, error) {
	baseClient, err := payment.NewBaseClient(app)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient,
	}, nil
}

// 图片上传API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter2_1_1.shtml
func (comp *Client) ApplyForBusiness(ctx context.Context, params *request.RequestMediaUpload) (*response.ResponseMediaUpload, error) {

	result := &response.ResponseMediaUpload{}

	var files *object.HashMap
	if params.File != "" {
		files = &object.HashMap{
			"media": params.File,
		}
	}

	var formData *kernel.UploadForm
	if params.Media != nil {
		formData = &kernel.UploadForm{
			Contents: []*kernel.UploadContent{
				&kernel.UploadContent{
					Name:  params.Media.Filename,
					Value: params.Media.Sha256,
				},
			},
		}
	}

	_, err := comp.BaseClient.HttpUpload(ctx, "cgi-bin/media/uploadimg", files, formData, nil, nil, &result)
	return result, err
}
