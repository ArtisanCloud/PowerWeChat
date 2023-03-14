package media

import (
	"context"
	"errors"
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/basicService/media/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"net/http"
	"os"
)

type Client struct {
	BaseClient *kernel.BaseClient

	AllowTypes []string
}

func NewClient(app kernel.ApplicationInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(&app, nil)
	if err != nil {
		return nil, err
	}
	client := &Client{
		BaseClient: baseClient,

		AllowTypes: []string{"image", "voice", "video", "thumb"},
	}

	client.BaseClient.BaseURI = "https://api.weixin.qq.com/cgi-bin/"
	return client, nil
}

// 新增临时素材
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/New_temporary_materials.html
func (comp *Client) UploadImage(ctx context.Context, path string) (*response.ResponseUploadMedia, error) {
	return comp.Upload(ctx, "image", path)
}

func (comp *Client) UploadVoice(ctx context.Context, path string) (*response.ResponseUploadMedia, error) {
	return comp.Upload(ctx, "voice", path)
}

func (comp *Client) UploadVideo(ctx context.Context, path string) (*response.ResponseUploadMedia, error) {
	return comp.Upload(ctx, "video", path)
}

func (comp *Client) UploadThumb(ctx context.Context, path string) (*response.ResponseUploadMedia, error) {
	return comp.Upload(ctx, "thumb", path)
}

// 上传临时素材
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/New_temporary_materials.html
func (comp *Client) Upload(ctx context.Context, mediaType string, path string) (*response.ResponseUploadMedia, error) {

	_, err := os.Stat(path)
	if (err != nil && os.IsExist(err)) && (err != nil && os.IsPermission(err)) {
		return nil, errors.New(fmt.Sprintf("File does not exist, or the file is unreadable: \"%s\"", path))
	}

	if !object.ContainsString(comp.AllowTypes, mediaType) {
		return nil, errors.New(fmt.Sprintf("Unsupported media type: '%s'", mediaType))
	}

	outResponse := &response.ResponseUploadMedia{}
	var files *object.HashMap
	if path != "" {
		files = &object.HashMap{
			"media": path,
		}
	} else {
		return nil, errors.New("path is empty")
	}

	_, err = comp.BaseClient.HttpUpload(ctx, "cgi-bin/media/upload", files, nil, &object.StringMap{
		"type": mediaType,
	}, nil, outResponse)

	return outResponse, err
}

// 获取临时素材
// https://work.weixin.qq.com/api/doc/90000/90135/90254
func (comp *Client) Get(ctx context.Context, mediaID string) (*http.Response, error) {

	header := &response.ResponseHeaderMedia{}
	response, err := comp.BaseClient.RequestRaw(ctx, "cgi-bin/media/get", http.MethodPost, &object.HashMap{
		"query": &object.StringMap{
			"media_id": mediaID,
		},
	}, header, nil)

	return response, err

}

// 获取高清语音素材
// https://work.weixin.qq.com/api/doc/90000/90135/90255
func (comp *Client) GetJSSDK(ctx context.Context, mediaID string) (*http.Response, error) {

	header := &response.ResponseHeaderMedia{}
	response, err := comp.BaseClient.RequestRaw(ctx, "cgi-bin/media/get/jssdk", http.MethodPost, &object.HashMap{
		"query": &object.StringMap{
			"media_id": mediaID,
		},
	}, header, nil)

	return response, err

}
