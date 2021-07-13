package media

import (
	"github.com/ArtisanCloud/go-libs/http/contract"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	response2 "github.com/ArtisanCloud/power-wechat/src/work/media/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

func (comp *Client) Get(mediaID string) (contract.ResponseContract, error) {

	//result := &response.ResponseWX{}
	result := ""
	header := &response2.ResponseHeaderMedia{}
	response := comp.RequestRaw("cgi-bin/media/get", "GET", &object.HashMap{
		"query": &object.StringMap{
			"media_id": mediaID,
		},
	}, header, &result)

	//fmt.Dump("Get:", header, result)
	return response.(contract.ResponseContract), nil

	//rs := rawResponse.(contract.ResponseContract)
	//header := (*rs.GetHeaders()).Header()
	//if header.Get("Content-Type") != "text/plain" {
	//
	//	config := *(*comp.App).GetContainer().Config
	//	var rs http.Response = http.Response{
	//		StatusCode: 200,
	//		Header:     nil,
	//	}
	//	postBodyBuf, _ := json.Marshal(rs)
	//	rs.Body = ioutil.NopCloser(bytes.NewBuffer(postBodyBuf))
	//
	//	returnResponse, err := comp.CastResponseToType(&rs, config["response_type"].(string))
	//	return returnResponse.(contract.ResponseContract), err
	//}

	//return rs, nil
}

func (comp *Client) UploadImage(path string, form *object.HashMap, outResponse interface{}) interface{} {
	files := &object.HashMap{
		"media": path,
	}
	return comp.HttpUpload("cgi-bin/media/uploadimg", files, form, nil, nil, outResponse)
}

func (comp *Client) UploadTempImage(path string, form *object.HashMap, outResponse interface{}) interface{} {
	return comp.Upload("image", path, form, outResponse)
}

func (comp *Client) UploadTempVoice(path string, form *object.HashMap, outResponse interface{}) interface{} {
	return comp.Upload("voice", path, form, outResponse)
}

func (comp *Client) UploadTempVideo(path string, form *object.HashMap, outResponse interface{}) interface{} {
	return comp.Upload("video", path, form, outResponse)
}

func (comp *Client) UploadTempFile(path string, form *object.HashMap, outResponse interface{}) interface{} {
	return comp.Upload("file", path, form, outResponse)
}

func (comp *Client) Upload(mediaType string, path string, form *object.HashMap, outResponse interface{}) interface{} {

	files := &object.HashMap{
		"media": path,
	}

	return comp.HttpUpload("cgi-bin/media/upload", files, form, &object.StringMap{
		"type": mediaType,
	}, nil, outResponse)
}
