package message

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-wechat/src/kernel"
	"github.com/ArtisanCloud/go-wechat/src/work/department/response"
)

type Client struct {
	*kernel.BaseClient
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/90205
func (comp *Client) Message(data *object.HashMap) *response.ResponseDepartmentCreate {

	result := &response.ResponseDepartmentCreate{}

	comp.HttpPostJson("cgi-bin/department/create", *data, nil, result)

	return result
}