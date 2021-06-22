package externalContact

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-wechat/src/kernel"
	response2 "github.com/ArtisanCloud/go-wechat/src/kernel/response"
	"github.com/ArtisanCloud/go-wechat/src/work/externalContact/response"
)

type MessageClient struct {
	*kernel.BaseClient
}

func NewMessageClient(app kernel.ApplicationInterface) *MessageClient {
	return &MessageClient{
		kernel.NewBaseClient(&app, nil),
	}
}

// 添加企业群发消息模板
// https://work.weixin.qq.com/api/doc#90000/90135/91560
func (comp *MessageClient) Submit(msg *object.HashMap) *response.ResponseAddMessageTemplate {

	result := &response.ResponseAddMessageTemplate{}

	params := comp.formatMessage(msg)

	comp.HttpPostJson("cgi-bin/externalcontact/add_msg_template", params, nil, result)

	return result
}

func (comp *MessageClient) Get(msgID string) *response2.ResponseWX {

	result := &response2.ResponseWX{}

	comp.HttpPostJson("cgi-bin/externalcontact/get_group_msg_result", &object.StringMap{
		"msgid": msgID,
	}, nil, result)

	return result
}

func (comp *MessageClient) SendWelcome(welcomeCode string, msg *object.HashMap) *response.ResponseAddMessageTemplate {

	result := &response.ResponseAddMessageTemplate{}

	formattedMsg := comp.formatMessage(msg)

	params := object.MergeHashMap(formattedMsg, &object.HashMap{
		"welcome_code": welcomeCode,
	})

	comp.HttpPostJson("cgi-bin/externalcontact/add_msg_template", params, nil, result)

	return result
}


func (comp *MessageClient) formatMessage(data *object.HashMap){
	params := data

	if params["text"] !=nil{
		params["text"] = comp.formatFields()
	}
}