package externalContact

import (
	"github.com/ArtisanCloud/go-wechat/src/kernel"
	"github.com/ArtisanCloud/go-wechat/src/work/externalContact/request"
	"github.com/ArtisanCloud/go-wechat/src/work/externalContact/response"
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
func (comp *ContactWayClient) Create( params *request.RequestAddContactWay) *response.ResponseAddContactWay {

	result := &response.ResponseAddContactWay{}

	comp.HttpGet("cgi-bin/externalcontact/add_contact_way", params, result)

	return result
}

