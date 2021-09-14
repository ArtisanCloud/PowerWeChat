package contactWay

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	response2 "github.com/ArtisanCloud/power-wechat/src/kernel/response"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/request"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

// 配置客户联系「联系我」方式.
// https://open.work.weixin.qq.com/api/doc/90000/90135/92572
func (comp *Client) Create(params *request.RequestAddContactWay) (*response.ResponseAddContactWay, error) {

	result := &response.ResponseAddContactWay{}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/add_contact_way", params, nil, nil, result)

	return result, err
}

func (comp *Client) Get(configID string) (*response.ResponseGetContactWay, error) {

	result := &response.ResponseGetContactWay{}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/get_contact_way", &object.StringMap{
		"config_id": configID,
	}, nil, nil, result)

	return result, err
}

func (comp *Client) Update(configID string, config *object.HashMap) (*response.ResponseAddContactWay, error) {

	result := &response.ResponseAddContactWay{}
	params := object.MergeHashMap(&object.HashMap{
		"config_id": configID,
	}, config)

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/update_contact_way", params, nil, nil, result)

	return result, err
}

func (comp *Client) Delete(configID string) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/del_contact_way", &object.StringMap{
		"config_id": configID,
	}, nil, nil, result)

	return result, err
}
