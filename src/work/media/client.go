package media

import (
	"github.com/ArtisanCloud/go-libs/http/contract"
	"github.com/ArtisanCloud/go-libs/http/response"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	response2 "github.com/ArtisanCloud/PowerWeChat/src/work/media/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

// 获取临时素材
// https://work.weixin.qq.com/api/doc/90000/90135/90254
func (comp *Client) Get(mediaID string) (contract.ResponseInterface, error) {

	result := ""
	header := &response2.ResponseHeaderMedia{}
	response, err := comp.RequestRaw("cgi-bin/media/get", "GET", &object.HashMap{
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
	header := &response2.ResponseHeaderMedia{}
	response, err := comp.RequestRaw("cgi-bin/media/get/jssdk", "GET", &object.HashMap{
		"query": &object.StringMap{
			"media_id": mediaID,
		},
	}, header, &result)

	return response.(contract.ResponseInterface), err

}

// 上传图片
// https://work.weixin.qq.com/api/doc/90000/90135/90256
func (comp *Client) UploadImage(path string, form *power.HashMap) (*response.HttpResponse, error) {
	files := &object.HashMap{
		"media": path,
	}
	rs, err := comp.HttpUpload("cgi-bin/media/uploadimg", files, form.ToHashMap(), nil, nil, nil)
	return rs.(*response.HttpResponse), err
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
// https://work.weixin.qq.com/api/doc/90000/90135/90253
func (comp *Client) Upload(mediaType string, path string, form *power.HashMap) (*response2.ResponseUploadMedia, error) {
	outResponse := &response2.ResponseUploadMedia{}
	files := &object.HashMap{
		"media": path,
	}

	_, err := comp.HttpUpload("cgi-bin/media/upload", files, form.ToHashMap(), &object.StringMap{
		"type": mediaType,
	}, nil, outResponse)

	return outResponse, err
}
