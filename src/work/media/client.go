package media

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/support"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/work/media/response"
	"github.com/pkg/errors"
	"net/http"
	"os"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(&app, nil)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient,
	}, nil
}

// 获取临时素材
// https://developer.work.weixin.qq.com/document/path/90254
func (comp *Client) Get(ctx context.Context, mediaID string) (*http.Response, error) {

	header := &response2.ResponseHeaderMedia{}
	response, err := comp.BaseClient.RequestRaw(ctx, "cgi-bin/media/get", http.MethodPost, &object.HashMap{
		"query": &object.StringMap{
			"media_id": mediaID,
		},
	}, header, nil)

	return response, err

}

// 获取高清语音素材
// https://developer.work.weixin.qq.com/document/path/90255
func (comp *Client) GetJSSDK(ctx context.Context, mediaID string) (*http.Response, error) {

	header := &response2.ResponseHeaderMedia{}
	response, err := comp.BaseClient.RequestRaw(ctx, "cgi-bin/media/get/jssdk", http.MethodPost, &object.HashMap{
		"query": &object.StringMap{
			"media_id": mediaID,
		},
	}, header, nil)

	return response, err

}

// 上传图片
// https://developer.work.weixin.qq.com/document/path/90256
func (comp *Client) UploadImage(ctx context.Context, path string, form *power.HashMap) (*response2.ResponseUploadImage, error) {
	result := &response2.ResponseUploadImage{}

	var files *object.HashMap
	if path != "" {
		files = &object.HashMap{
			"media": path,
		}
	}

	var formData *kernel.UploadForm
	if form != nil {
		formData = &kernel.UploadForm{
			Contents: []*kernel.UploadContent{
				&kernel.UploadContent{
					Name:  (*form)["name"].(string),
					Value: (*form)["value"],
				},
			},
		}
	}

	_, err := comp.BaseClient.HttpUpload(ctx, "cgi-bin/media/uploadimg", files, formData, nil, nil, &result)
	return result, err
}

func (comp *Client) UploadTempImageUrl(ctx context.Context, url string, form *power.HashMap) (*response2.ResponseUploadMedia, error) {
	path, err := support.DownloadFileFromURL(url)
	if err != nil {
		return nil, errors.Wrap(err, "failed to download file from URL")
	}
	defer os.Remove(path)
	return comp.Upload(ctx, "image", path, form)
}

func (comp *Client) UploadTempVoiceUrl(ctx context.Context, url string, form *power.HashMap) (*response2.ResponseUploadMedia, error) {
	path, err := support.DownloadFileFromURL(url)
	if err != nil {
		return nil, errors.Wrap(err, "failed to download file from URL")
	}
	defer os.Remove(path)
	return comp.Upload(ctx, "voice", path, form)
}

func (comp *Client) UploadTempVideoUrl(ctx context.Context, url string, form *power.HashMap) (*response2.ResponseUploadMedia, error) {
	path, err := support.DownloadFileFromURL(url)
	if err != nil {
		return nil, errors.Wrap(err, "failed to download file from URL")
	}
	defer os.Remove(path)
	return comp.Upload(ctx, "video", path, form)
}

func (comp *Client) UploadTempFileUrl(ctx context.Context, url string, form *power.HashMap) (*response2.ResponseUploadMedia, error) {
	path, err := support.DownloadFileFromURL(url)
	if err != nil {
		return nil, errors.Wrap(err, "failed to download file from URL")
	}
	defer os.Remove(path)
	return comp.Upload(ctx, "file", path, form)
}

func (comp *Client) UploadTempImage(ctx context.Context, path string, form *power.HashMap) (*response2.ResponseUploadMedia, error) {
	return comp.Upload(ctx, "image", path, form)
}

func (comp *Client) UploadTempVoice(ctx context.Context, path string, form *power.HashMap) (*response2.ResponseUploadMedia, error) {
	return comp.Upload(ctx, "voice", path, form)
}

func (comp *Client) UploadTempVideo(ctx context.Context, path string, form *power.HashMap) (*response2.ResponseUploadMedia, error) {
	return comp.Upload(ctx, "video", path, form)
}

func (comp *Client) UploadTempFile(ctx context.Context, path string, form *power.HashMap) (*response2.ResponseUploadMedia, error) {
	return comp.Upload(ctx, "file", path, form)
}

// 上传临时素材
// https://developer.work.weixin.qq.com/document/path/90253
func (comp *Client) Upload(ctx context.Context, mediaType string, path string, form *power.HashMap) (*response2.ResponseUploadMedia, error) {
	outResponse := &response2.ResponseUploadMedia{}

	var files *object.HashMap
	if path != "" {
		files = &object.HashMap{
			"media": path,
		}
	}

	var formData *kernel.UploadForm
	if form != nil {
		formData = &kernel.UploadForm{
			Contents: []*kernel.UploadContent{
				&kernel.UploadContent{
					Name:  (*form)["name"].(string),
					Value: (*form)["value"],
				},
			},
		}
	}

	_, err := comp.BaseClient.HttpUpload(ctx, "cgi-bin/media/upload", files, formData, &object.StringMap{
		"type": mediaType,
	}, nil, outResponse)

	return outResponse, err
}
