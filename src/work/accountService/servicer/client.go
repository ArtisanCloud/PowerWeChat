package servicer

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/work/accountService/servicer/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

// 添加接待人员
// https://work.weixin.qq.com/api/doc/90000/90135/94646
func (comp *Client) Add(openKFID string, userIDList string) (*response.ResponseServicerAdd, error) {

	result := &response.ResponseServicerAdd{}

	options := &object.HashMap{
		"open_kfid":     openKFID,
		"userid_list": userIDList,
	}

	_, err := comp.HttpPostJson("cgi-bin/kf/servicer/add", options, nil, nil, result)

	return result, err
}

// 删除接待人员
// https://work.weixin.qq.com/api/doc/90000/90135/94647
func (comp *Client) Del(openKFID string, userIDList string) (*response.ResponseServicerDel, error) {

	result := &response.ResponseServicerDel{}

	options := &object.HashMap{
		"open_kfid":     openKFID,
		"userid_list": userIDList,
	}

	_, err := comp.HttpPostJson("cgi-bin/kf/servicer/del", options, nil, nil, result)

	return result, err
}


// 获取接待人员列表
// https://work.weixin.qq.com/api/doc/90000/90135/94645
func (comp *Client) List(openKFID string) (*response.ResponseServicerList, error) {

	result := &response.ResponseServicerList{}

	params := &object.StringMap{
		"open_kfid":     openKFID,
	}

	_, err := comp.HttpPostJson("cgi-bin/kf/servicer/list", nil, params, nil, result)

	return result, err
}