package media

import (
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/work/media/response"
	"net/http"
)

type Client struct {
	*kernel.BaseClient
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
func (comp *Client) Get(mediaID string) (*http.Response, error) {

	result := ""
	header := &response2.ResponseHeaderMedia{}
	response, err := comp.RequestRaw("cgi-bin/media/get", "GET", &object.HashMap{
		"query": &object.StringMap{
			"media_id": mediaID,
		},
	}, header, &result)

	return response, err

}

// 获取高清语音素材
// https://developer.work.weixin.qq.com/document/path/90255
func (comp *Client) GetJSSDK(mediaID string) (*http.Response, error) {

	result := ""
	header := &response2.ResponseHeaderMedia{}
	response, err := comp.RequestRaw("cgi-bin/media/get/jssdk", "GET", &object.HashMap{
		"query": &object.StringMap{
			"media_id": mediaID,
		},
	}, header, &result)

	return response, err

}

// 上传图片
// https://developer.work.weixin.qq.com/document/path/90256
func (comp *Client) UploadImage(path string, form *power.HashMap) (*response2.ResponseUploadImage, error) {
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

	_, err := comp.HttpUpload("cgi-bin/media/uploadimg", files, formData, nil, nil, &result)
	return result, err
}

func (comp *Client) UploadTempImage(path string, form *power.HashMap) (*response2.ResponseUploadMedia, error) {
	return comp.Upload("image", path, form)
}

func (comp *Client) UploadTempVoice(path string, form *power.HashMap) (*response2.ResponseUploadMedia, error) {
	return comp.Upload("voice", path, form)
}

func (comp *Client) UploadTempVideo(path string, form *power.HashMap) (*response2.ResponseUploadMedia, error) {
	return comp.Upload("video", path, form)
}

func (comp *Client) UploadTempFile(path string, form *power.HashMap) (*response2.ResponseUploadMedia, error) {
	return comp.Upload("file", path, form)
}

// 上传临时素材
// https://developer.work.weixin.qq.com/document/path/90253
func (comp *Client) Upload(mediaType string, path string, form *power.HashMap) (*response2.ResponseUploadMedia, error) {
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

	_, err := comp.HttpUpload("cgi-bin/media/upload", files, formData, &object.StringMap{
		"type": mediaType,
	}, nil, outResponse)

	return outResponse, err
}
