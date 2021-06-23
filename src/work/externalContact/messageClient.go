package externalContact

import (
	"errors"
	"fmt"
	fmt2 "github.com/ArtisanCloud/go-libs/fmt"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-wechat/src/kernel"
	response2 "github.com/ArtisanCloud/go-wechat/src/kernel/response"
	"github.com/ArtisanCloud/go-wechat/src/work/externalContact/request"
	"github.com/ArtisanCloud/go-wechat/src/work/externalContact/response"
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
	fmt2.Dump(params)
	//comp.HttpPostJson("cgi-bin/externalcontact/add_msg_template", params, nil, result)

	return result, nil
}

func (comp *MessageClient) Get(msgID string) *response2.ResponseWX {

	result := &response2.ResponseWX{}

	comp.HttpPostJson("cgi-bin/externalcontact/get_group_msg_result", &object.StringMap{
		"msgid": msgID,
	}, nil, result)

	return result
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

	comp.HttpPostJson("cgi-bin/externalcontact/add_msg_template", params, nil, result)

	return result, nil
}

func (comp *MessageClient) formatMessage(data *object.HashMap) (*object.HashMap, error) {
	params := *data

	if params["text"] != nil {

		switch params["text"].(type){
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
				_, err := comp.formatFields(obj.(request.ImageOfMessage))
				if err != nil {
					return nil, err
				}

				break
			case request.LinkOfMessage:
				_, err := comp.formatFields(params["link"].(request.LinkOfMessage))
				if err != nil {
					return nil, err
				}

				break
			case request.MiniProgramOfMessage:
				_, err := comp.formatFields(params["miniprogram"].(request.MiniProgramOfMessage))
				if err != nil {
					return nil, err
				}

				break
			case request.VideoOfMessage:
				_, err := comp.formatFields(params["video"].(request.VideoOfMessage))
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
	inputDataType := reflect.TypeOf(data)

	// new a return data of default type
	dataReturn := reflect.New(inputDataType)

	for i := 0; i < dataReturn.NumField(); i++ {
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
