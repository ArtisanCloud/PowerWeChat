package messageTemplate

import (
	"errors"
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/src/work/externalContact/messageTemplate/request"
	"github.com/ArtisanCloud/PowerWeChat/src/work/externalContact/messageTemplate/response"
	"reflect"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

var required = []string{"content", "title", "url", "pic_media_id", "appid", "page"}

// 创建企业群发
// https://work.weixin.qq.com/api/doc/90000/90135/92135
func (comp *Client) AddMsgTemplate(options *request.RequestAddMsgTemplate) (*response.ResponseAddMessageTemplate, error) {

	result := &response.ResponseAddMessageTemplate{}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/add_msg_template", options, nil, nil, result)

	return result, err
}

// 获取群发记录列表
// https://work.weixin.qq.com/api/doc/90000/90135/93338
func (comp *Client) GetGroupMsgListV2(options *request.RequestGetGroupMsgListV2) (*response.ResponseGetGroupMsgListV2, error) {

	result := &response.ResponseGetGroupMsgListV2{}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/get_groupmsg_list_v2", options, nil, nil, result)

	return result, err
}

// 获取群发成员发送任务列表
// https://work.weixin.qq.com/api/doc/90000/90135/93338
func (comp *Client) GetGroupMsgTask(msgID string, limit int, cursor string) (*response.ResponseGetGroupMsgTask, error) {

	result := &response.ResponseGetGroupMsgTask{}
	options := &object.HashMap{
		"msgid":       msgID,
		"limit":       limit,
		"msgcursorid": cursor,
	}
	_, err := comp.HttpPostJson("cgi-bin/externalcontact/get_groupmsg_task", options, nil, nil, result)

	return result, err
}

// 获取企业群发成员执行结果
// https://work.weixin.qq.com/api/doc/90000/90135/93338
func (comp *Client) GetGroupMsgSendResult(msgID string, userID string, limit int, cursor string) (*response.ResponseGetGroupMsgSendResult, error) {

	result := &response.ResponseGetGroupMsgSendResult{}
	options := &object.HashMap{
		"msgid":       msgID,
		"userid":      userID,
		"limit":       limit,
		"msgcursorid": cursor,
	}
	_, err := comp.HttpPostJson("cgi-bin/externalcontact/get_groupmsg_send_result", options, nil, nil, result)

	return result, err
}

// 发送新客户欢迎语
// https://work.weixin.qq.com/api/doc/90000/90135/93338
func (comp *Client) SendWelcomeMsg(options *request.RequestSendWelcomeMsg) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/send_welcome_msg", options, nil, nil, result)

	return result, err
}

func (comp *Client) formatMessage(data *object.HashMap) (*object.HashMap, error) {
	params := *data

	if params["text"] != nil {

		switch params["text"].(type) {
		case request.TextOfMessage:
			_, err := comp.formatFields(params["text"].(request.TextOfMessage))
			if err != nil {
				return nil, err
			}
		default:
		}
	}

	if params["attachments"] != nil {
		attachements := params["attachments"].([]interface{})

		for _, obj := range attachements {

			switch obj.(type) {
			case request.ImageOfMessage:
				_, err := comp.formatFields(obj.(request.ImageOfMessage).Image)
				if err != nil {
					return nil, err
				}

				break
			case request.LinkOfMessage:
				_, err := comp.formatFields(obj.(request.LinkOfMessage).Link)
				if err != nil {
					return nil, err
				}

				break
			case request.MiniProgramOfMessage:
				_, err := comp.formatFields(obj.(request.MiniProgramOfMessage).MiniProgram)
				if err != nil {
					return nil, err
				}

				break
			case request.VideoOfMessage:
				_, err := comp.formatFields(obj.(request.VideoOfMessage).Video)
				if err != nil {
					return nil, err
				}

				break
			default:
			}
		}
	}

	return &params, nil
}

func (comp *Client) formatFields(data interface{}) (interface{}, error) {

	// input data
	inputValues := reflect.ValueOf(data)

	// get type of default Data
	inputDataType := inputValues.Type()

	// new a return data of default type
	dataReturn := reflect.New(inputDataType)

	for i := 0; i < inputValues.NumField(); i++ {
		key := inputDataType.Field(i).Name
		value := inputValues.Field(i).Interface()

		fmt.Printf("Field: %s\tValue: %v\n", key, value)
		if object.InArray(key, required) && value == nil {
			return nil, errors.New(fmt.Sprintf("Attribute \"%s\" can not be empty!", key))
		}

		dataReturn.Elem().Field(i).SetString(value.(string))

	}

	return dataReturn, nil

}
