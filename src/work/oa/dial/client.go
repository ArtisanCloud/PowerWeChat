package dial

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/work/oa/dial/request"
	"github.com/ArtisanCloud/power-wechat/src/work/oa/dial/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

// 获取公费电话拨打记录
// https://work.weixin.qq.com/api/doc/90000/90135/93662
func (comp *Client) GetDialRecord(options *request.RequestDialGetDialRecord) (*response.ResponseDialGetDialRecord, error) {

	result := &response.ResponseDialGetDialRecord{}

	_, err := comp.HttpPostJson("cgi-bin/dial/get_dial_record", options, nil, nil, result)

	return result, err
}
