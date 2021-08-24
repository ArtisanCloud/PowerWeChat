package externalContact

import (
	"errors"
	"fmt"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/request"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/response"
	"reflect"
)

type MessageClient struct {
	*kernel.BaseClient
}

func NewMessageClient(app kernel.ApplicationInterface) *MessageClient {
	return &MessageClient{
		kernel.NewBaseClient(&app, nil),
	}
}

var required = []string{"content", "title", "url", "pic_media_id", "appid", "page"}

// 添加企业群发消息模板
// https://work.weixin.qq.com/api/doc#90000/90135/91560
func (comp *MessageClient) Submit(msg *object.HashMap) (*response.ResponseAddMessageTemplate, error) {

	result := &response.ResponseAddMessageTemplate{}

	params, err := comp.formatMessage(msg)
	if err != nil {
		return nil, err
	}
	_, err = comp.HttpPostJson("cgi-bin/externalcontact/add_msg_template", params, nil, nil, result)

	return result, err
}

func (comp *MessageClient) Get(msgID string) (*response.ResponseGetGroupMesResult, error) {

	result := &response.ResponseGetGroupMesResult{}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/get_group_msg_result", &object.StringMap{
		"msgid": msgID,
	}, nil, nil, result)

	return result, err
}

func (comp *MessageClient) SendWelcome(welcomeCode string, msg *object.HashMap) (*response.ResponseAddMessageTemplate, error) {

	result := &response.ResponseAddMessageTemplate{}

	formattedMsg, err := comp.formatMessage(msg)
	if err != nil {
		return nil, err
	}

	params := object.MergeHashMap(formattedMsg, &object.HashMap{
		"welcome_code": welcomeCode,
	})

	_, err = comp.HttpPostJson("cgi-bin/externalcontact/send_welcome_msg", params, nil, nil, result)

	return result, err
}

func (comp *MessageClient) formatMessage(data *object.HashMap) (*object.HashMap, error) {
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

func (comp *MessageClient) formatFields(data interface{}) (interface{}, error) {

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
