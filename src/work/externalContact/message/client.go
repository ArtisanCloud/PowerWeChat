package message

import (
	"errors"
	"fmt"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	request2 "github.com/ArtisanCloud/power-wechat/src/work/externalContact/message/request"
	response2 "github.com/ArtisanCloud/power-wechat/src/work/externalContact/messageTemplate/response"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/response"
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

// 添加企业群发消息模板
// https://work.weixin.qq.com/api/doc#90000/90135/91560
func (comp  *Client) Submit(msg *object.HashMap) (*response2.ResponseAddMessageTemplate, error) {

	result := &response2.ResponseAddMessageTemplate{}

	params, err := comp.formatMessage(msg)
	if err != nil {
		return nil, err
	}
	_, err = comp.HttpPostJson("cgi-bin/externalcontact/add_msg_template", params, nil, nil, result)

	return result, err
}

func (comp  *Client) Get(msgID string) (*response.ResponseGetGroupMesResult, error) {

	result := &response.ResponseGetGroupMesResult{}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/get_group_msg_result", &object.StringMap{
		"msgid": msgID,
	}, nil, nil, result)

	return result, err
}

func (comp  *Client) SendWelcome(welcomeCode string, msg *power.HashMap) (*response2.ResponseAddMessageTemplate, error) {

	result := &response2.ResponseAddMessageTemplate{}

	formattedMsg, err := comp.formatMessage(msg.ToHashMap())
	if err != nil {
		return nil, err
	}

	params := object.MergeHashMap(formattedMsg, &object.HashMap{
		"welcome_code": welcomeCode,
	})

	_, err = comp.HttpPostJson("cgi-bin/externalcontact/send_welcome_msg", params, nil, nil, result)

	return result, err
}

func (comp  *Client) formatMessage(data *object.HashMap) (*object.HashMap, error) {
	params := *data

	if params["text"] != nil {

		switch params["text"].(type) {
		case request2.TextOfMessage:
			_, err := comp.formatFields(params["text"].(request2.TextOfMessage))
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
			case request2.ImageOfMessage:
				_, err := comp.formatFields(obj.(request2.ImageOfMessage).Image)
				if err != nil {
					return nil, err
				}

				break
			case request2.LinkOfMessage:
				_, err := comp.formatFields(obj.(request2.LinkOfMessage).Link)
				if err != nil {
					return nil, err
				}

				break
			case request2.MiniProgramOfMessage:
				_, err := comp.formatFields(obj.(request2.MiniProgramOfMessage).MiniProgram)
				if err != nil {
					return nil, err
				}

				break
			case request2.VideoOfMessage:
				_, err := comp.formatFields(obj.(request2.VideoOfMessage).Video)
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

func (comp  *Client) formatFields(data interface{}) (interface{}, error) {

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
