package department

import (
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/work/department/request"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/work/department/response"
)

type Client struct {
	*kernel.BaseClient
}

// 创建部门
// https://open.work.weixin.qq.com/api/doc/90000/90135/90205
func (comp *Client) Create(params *request.RequestDepartmentUpsert) (*response.ResponseDepartmentCreate, error) {

	result := &response.ResponseDepartmentCreate{}

	_, err := comp.HttpPostJson("cgi-bin/department/create", params, nil, nil, result)

	return result, err
}

// 更新部门
// https://open.work.weixin.qq.com/api/doc/90000/90135/90206
func (comp *Client) Update(params *request.RequestDepartmentUpsert) (*response.ResponseDepartmentUpdate, error) {

	result := &response.ResponseDepartmentUpdate{}

	_, err := comp.HttpPostJson("cgi-bin/department/update", params, nil, nil, result)

	return result, err
}

// 删除部门
// https://open.work.weixin.qq.com/api/doc/90000/90135/90207
func (comp *Client) Delete(id int) (*response.ResponseDepartmentDelete, error) {

	result := &response.ResponseDepartmentDelete{}

	_, err := comp.HttpGet("cgi-bin/department/delete", &object.StringMap{
		"id": fmt.Sprintf("%d", id),
	}, nil, result)

	return result, err
}

// 获取部门列表
// https://open.work.weixin.qq.com/api/doc/90000/90135/90208
func (comp *Client) List(id int) (*response.ResponseDepartmentList, error) {

	result := &response.ResponseDepartmentList{}

	_, err := comp.HttpGet("cgi-bin/department/list", &object.StringMap{
		"id": fmt.Sprintf("%d", id),
	}, nil, result)

	return result, err
}

// 获取子部门ID列表
// https://developer.work.weixin.qq.com/document/path/95350
func (comp *Client) SimpleList(id int) (*response.ResponseDepartmentIDList, error) {

	result := &response.ResponseDepartmentIDList{}

	_, err := comp.HttpGet("cgi-bin/department/simplelist", &object.StringMap{
		"id": fmt.Sprintf("%d", id),
	}, nil, result)

	return result, err
}
