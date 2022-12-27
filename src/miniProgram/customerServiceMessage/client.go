package customerServiceMessage

import (
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/customerServiceMessage/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/customerServiceMessage/response"
	response4 "github.com/ArtisanCloud/PowerWeChat/v3/src/work/media/response"
	"net/http"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// 获取客服消息内的临时素材
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.getTempMedia.html
func (comp *Client) GetTempMedia(mediaID string) (*http.Response, error) {

	var result string
	var header = &response4.ResponseHeaderMedia{}

	params := &object.HashMap{
		"query": &object.StringMap{
			"media_id": mediaID,
		},
	}

	rs, err := comp.BaseClient.RequestRaw("cgi-bin/media/get", "GET", params, &header, &result)

	return rs, err
}

// 发送客服消息给用户
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.send.html
func (comp *Client) Send(toUser string, msgType string, msg interface{}) (*response2.ResponseMiniProgram, error) {

	result := &response2.ResponseMiniProgram{}

	data := &object.HashMap{
		"touser":  toUser,
		"msgtype": msgType,
		msgType:   msg,
	}

	_, err := comp.BaseClient.HttpPostJson("cgi-bin/message/custom/send", data, nil, nil, result)

	return result, err
}
func (comp *Client) SendText(toUser string, msg *request.CustomerServiceMsgText) (*response2.ResponseMiniProgram, error) {
	return comp.Send(toUser, "text", msg)
}
func (comp *Client) SendImage(toUser string, msg *request.CustomerServiceMsgImage) (*response2.ResponseMiniProgram, error) {
	return comp.Send(toUser, "image", msg)
}
func (comp *Client) SendLink(toUser string, msg *request.CustomerServiceMsgLink) (*response2.ResponseMiniProgram, error) {
	return comp.Send(toUser, "link", msg)
}
func (comp *Client) SendMiniProgramPage(toUser string, msg *request.CustomerServiceMsgMpPage) (*response2.ResponseMiniProgram, error) {
	return comp.Send(toUser, "miniprogrampage", msg)
}

// 下发客服当前输入状态给用户
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.setTyping.html
func (comp *Client) SetTyping(toUser string, command string) (*response2.ResponseMiniProgram, error) {

	result := &response2.ResponseMiniProgram{}

	data := &object.HashMap{
		"touser":  toUser,
		"command": command,
	}

	_, err := comp.BaseClient.HttpPostJson("cgi-bin/message/custom/typing", data, nil, nil, result)

	return result, err
}

// 把媒体文件上传到微信服务器
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.uploadTempMedia.html
func (comp *Client) UploadTempMedia(mediaType string, path string, form *power.HashMap) (*response.ResponseCustomerServiceMessageUploadTempMedia, error) {

	result := &response.ResponseCustomerServiceMessageUploadTempMedia{}

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

	_, err := comp.BaseClient.HttpUpload("cgi-bin/media/upload", files, formData, &object.StringMap{
		"type": mediaType,
	}, nil, result)

	return result, err
}
