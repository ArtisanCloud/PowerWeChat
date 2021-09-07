package customerServiceMessage

import (
	response3 "github.com/ArtisanCloud/go-libs/http/response"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	response2 "github.com/ArtisanCloud/power-wechat/src/kernel/response"
	"github.com/ArtisanCloud/power-wechat/src/miniProgram/customerServiceMessage/response"
	response4 "github.com/ArtisanCloud/power-wechat/src/work/media/response"
	"net/http"
)

type Client struct {
	*kernel.BaseClient
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

	rs, err := comp.RequestRaw("cgi-bin/media/get", "GET", params, &header, &result)

	httpRS := rs.(*response3.HttpResponse).Response

	return httpRS, err
}

// 发送客服消息给用户
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.send.html
func (comp *Client) Send(toUser string, msgType string, text *power.HashMap) (*response2.ResponseMiniProgram, error) {

	result := &response2.ResponseMiniProgram{}

	data := &object.HashMap{
		"touser":  toUser,
		"msgtype": msgType,
		"text":    text,
	}

	_, err := comp.HttpPostJson("cgi-bin/message/custom/send", data, nil, nil, result)

	return result, err
}

// 下发客服当前输入状态给用户
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.setTyping.html
func (comp *Client) SetTyping(toUser string, command string) (*response2.ResponseMiniProgram, error) {

	result := &response2.ResponseMiniProgram{}

	data := &object.HashMap{
		"touser":  toUser,
		"command": command,
	}

	_, err := comp.HttpPostJson("cgi-bin/message/custom/typing", data, nil, nil, result)

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

	_, err := comp.HttpUpload("cgi-bin/media/upload", files, form.ToHashMap(), &object.StringMap{
		"type": mediaType,
	}, nil, result)

	return result, err
}
