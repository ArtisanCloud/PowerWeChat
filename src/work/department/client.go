package department

import (
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/work/department/response"
)

type Client struct {
	*kernel.BaseClient
}

// 创建部门
// https://open.work.weixin.qq.com/api/doc/90000/90135/90205
func (comp *Client) Create(data *power.HashMap) (*response.ResponseDepartmentCreate, error) {

	result := &response.ResponseDepartmentCreate{}

	_, err := comp.HttpPostJson("cgi-bin/department/create", data, nil, nil, result)

	return result, err
}

// 更新部门
// https://open.work.weixin.qq.com/api/doc/90000/90135/90206
func (comp *Client) Update(id int, data *power.HashMap) (*response.ResponseDepartmentUpdate, error) {

	result := &response.ResponseDepartmentUpdate{}

	(*data)["id"] = id
	_, err := comp.HttpPostJson("cgi-bin/department/update", data, nil, nil, result)

	return result, err
}

// 删除部门
// https://open.work.weixin.qq.com/api/doc/90000/90135/90207
func (comp *Client) Delete(id int) (*response.ResponseDepartmentDelete, error) {

	result := &response.ResponseDepartmentDelete{}

	_, err := comp.HttpGet("cgi-bin/department/delete", &object.HashMap{"id": id}, nil, result)

	return result, err
}

// 获取部门列表
// https://open.work.weixin.qq.com/api/doc/90000/90135/90208
func (comp *Client) List() (*response.ResponseDepartmentList, error) {

	result := &response.ResponseDepartmentList{}

	_, err := comp.HttpGet("cgi-bin/department/list", nil, nil, result)

	return result, err
}
