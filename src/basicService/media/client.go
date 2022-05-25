package media

import (
	"errors"
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/http/contract"
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/basicService/media/response"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	"os"
)

type Client struct {
	*kernel.BaseClient

	AllowTypes []string
}

func NewClient(app kernel.ApplicationInterface) *Client {
	client := &Client{
		BaseClient: kernel.NewBaseClient(&app, nil),

		AllowTypes: []string{"image", "voice", "video", "thumb"},
	}

	client.BaseClient.HttpRequest.BaseURI = "https://api.weixin.qq.com/cgi-bin/"
	return client
}

// 新增临时素材
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/New_temporary_materials.html
func (comp *Client) UploadImage(path string) (*response.ResponseUploadMedia, error) {
	return comp.Upload("image", path)
}

func (comp *Client) UploadVoice(path string) (*response.ResponseUploadMedia, error) {
	return comp.Upload("voice", path)
}

func (comp *Client) UploadVideo(path string) (*response.ResponseUploadMedia, error) {
	return comp.Upload("video", path)
}

func (comp *Client) UploadThumb(path string) (*response.ResponseUploadMedia, error) {
	return comp.Upload("thumb", path)
}

// 上传临时素材
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/New_temporary_materials.html
func (comp *Client) Upload(mediaType string, path string) (*response.ResponseUploadMedia, error) {

	_, err := os.Stat(path)
	if (err != nil && os.IsExist(err)) && (err != nil && os.IsPermission(err)) {
		return nil, errors.New(fmt.Sprintf("File does not exist, or the file is unreadable: \"%s\"", path))
	}

	if !object.ContainsString(comp.AllowTypes, mediaType) {
		return nil, errors.New(fmt.Sprintf("Unsupported media type: '%s'", mediaType))
	}

	outResponse := &response.ResponseUploadMedia{}
	files := &object.HashMap{
		"media": path,
	}

	_, err = comp.HttpUpload("media/upload", files, nil, &object.StringMap{
		"type": mediaType,
	}, nil, outResponse)

	return outResponse, err
}

// 获取临时素材
// https://work.weixin.qq.com/api/doc/90000/90135/90254
func (comp *Client) Get(mediaID string) (contract.ResponseInterface, error) {

	result := ""
	header := &response.ResponseHeaderMedia{}
	response, err := comp.RequestRaw("media/get", "GET", &object.HashMap{
		"query": &object.StringMap{
			"media_id": mediaID,
		},
	}, header, &result)

	return response.(contract.ResponseInterface), err

}

// 获取高清语音素材
// https://work.weixin.qq.com/api/doc/90000/90135/90255
func (comp *Client) GetJSSDK(mediaID string) (contract.ResponseInterface, error) {

	result := ""
	header := &response.ResponseHeaderMedia{}
	response, err := comp.RequestRaw("media/get/jssdk", "GET", &object.HashMap{
		"query": &object.StringMap{
			"media_id": mediaID,
		},
	}, header, &result)

	return response.(contract.ResponseInterface), err

}
