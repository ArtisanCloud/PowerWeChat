package department

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-wechat/src/kernel"
	"github.com/ArtisanCloud/go-wechat/src/work/department/response"
	"strconv"
)

type Client struct {
	*kernel.BaseClient
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/90205
func (comp *Client) Create(data *object.HashMap) *response.ResponseDepartmentCreate {

	result := &response.ResponseDepartmentCreate{}

	comp.HttpPostJson("cgi-bin/department/create", data, nil, result)

	return result
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/90206
func (comp *Client) Update(id int, data *object.HashMap) *response.ResponseDepartmentUpdate {

	result := &response.ResponseDepartmentUpdate{}

	(*data)["id"] = id
	comp.HttpPostJson("cgi-bin/department/update", data, nil, result)

	return result
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/90207
func (comp *Client) Delete(id int) *response.ResponseDepartmentDelete {

	result := &response.ResponseDepartmentDelete{}

	comp.HttpGet("cgi-bin/department/delete", object.StringMap{"id":strconv.Itoa(id)}, result)

	return result
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/90208
func (comp *Client) List() *response.ResponseDepartmentList {

	result := &response.ResponseDepartmentList{}

	comp.HttpGet("cgi-bin/department/list", nil, result)

	return result
}
