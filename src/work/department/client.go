package department

import (
	"context"
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/department/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/department/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// 创建部门
// https://developer.work.weixin.qq.com/document/path/90205
func (comp *Client) Create(ctx context.Context, params *request.RequestDepartmentInsert) (*response.ResponseDepartmentCreate, error) {

	result := &response.ResponseDepartmentCreate{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/department/create", params, nil, nil, result)

	return result, err
}

// 更新部门
// https://developer.work.weixin.qq.com/document/path/90206
func (comp *Client) Update(ctx context.Context, params *request.RequestDepartmentUpdate) (*response.ResponseDepartmentUpdate, error) {

	result := &response.ResponseDepartmentUpdate{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/department/update", params, nil, nil, result)

	return result, err
}

// 删除部门
// https://developer.work.weixin.qq.com/document/path/90207
func (comp *Client) Delete(ctx context.Context, id int) (*response.ResponseDepartmentDelete, error) {

	result := &response.ResponseDepartmentDelete{}

	_, err := comp.BaseClient.HttpGet(ctx, "cgi-bin/department/delete", &object.StringMap{
		"id": fmt.Sprintf("%d", id),
	}, nil, result)

	return result, err
}

// 获取部门列表
// https://developer.work.weixin.qq.com/document/path/90208
func (comp *Client) List(ctx context.Context, id int) (*response.ResponseDepartmentList, error) {

	result := &response.ResponseDepartmentList{}

	_, err := comp.BaseClient.HttpGet(ctx, "cgi-bin/department/list", &object.StringMap{
		"id": fmt.Sprintf("%d", id),
	}, nil, result)

	return result, err
}

// 获取子部门ID列表
// https://developer.work.weixin.qq.com/document/path/95350
func (comp *Client) SimpleList(ctx context.Context, id int) (*response.ResponseDepartmentIDList, error) {

	result := &response.ResponseDepartmentIDList{}

	_, err := comp.BaseClient.HttpGet(ctx, "cgi-bin/department/simplelist", &object.StringMap{
		"id": fmt.Sprintf("%d", id),
	}, nil, result)

	return result, err
}

// 获取单个部门详情
// https://developer.work.weixin.qq.com/document/path/95351
func (comp *Client) Get(ctx context.Context, id int) (*response.ResponseDepartmentGet, error) {

	result := &response.ResponseDepartmentGet{}

	_, err := comp.BaseClient.HttpGet(ctx, "cgi-bin/department/get", &object.StringMap{
		"id": fmt.Sprintf("%d", id),
	}, nil, result)

	return result, err
}
