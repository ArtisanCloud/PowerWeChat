package externalContact

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	response2 "github.com/ArtisanCloud/power-wechat/src/kernel/response"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/request"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/response"
)

type ContactWayClient struct {
	*kernel.BaseClient
}

func NewContactClient(app kernel.ApplicationInterface) *ContactWayClient {
	return &ContactWayClient{
		kernel.NewBaseClient(&app, nil),
	}
}

// 配置客户联系「联系我」方式.
// https://open.work.weixin.qq.com/api/doc/90000/90135/92572
func (comp *ContactWayClient) Create(params *request.RequestAddContactWay) *response.ResponseAddContactWay {

	result := &response.ResponseAddContactWay{}

	comp.HttpPostJson("cgi-bin/externalcontact/add_contact_way", params, nil, result)

	return result
}

func (comp *ContactWayClient) Get(configID string) *response.ResponseGetContactWay {

	result := &response.ResponseGetContactWay{}

	comp.HttpPostJson("cgi-bin/externalcontact/get_contact_way", &object.StringMap{
		"config_id": configID,
	},nil, result)

	return result
}

func (comp *ContactWayClient) Update(configID string, config *object.HashMap) *response.ResponseAddContactWay {

	result := &response.ResponseAddContactWay{}
	params := object.MergeHashMap(&object.HashMap{
		"config_id": configID,
	}, config)

	comp.HttpPostJson("cgi-bin/externalcontact/update_contact_way", params, nil, result)

	return result
}

func (comp *ContactWayClient) Delete(configID string) *response2.ResponseWX {

	result := &response2.ResponseWX{}

	comp.HttpPostJson("cgi-bin/externalcontact/del_contact_way", &object.StringMap{
		"config_id": configID,
	}, nil, result)

	return result
}
