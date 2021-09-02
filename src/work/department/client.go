package department

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/work/department/response"
	"strconv"
)

type Client struct {
	*kernel.BaseClient
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/90205
func (comp *Client) Create(data *power.HashMap) (*response.ResponseDepartmentCreate, error) {

	result := &response.ResponseDepartmentCreate{}

	_, err := comp.HttpPostJson("cgi-bin/department/create", data, nil, nil, result)

	return result, err
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/90206
func (comp *Client) Update(id int, data *power.HashMap) (*response.ResponseDepartmentUpdate, error) {

	result := &response.ResponseDepartmentUpdate{}

	(*data)["id"] = id
	_, err := comp.HttpPostJson("cgi-bin/department/update", data, nil, nil, result)

	return result, err
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/90207
func (comp *Client) Delete(id int) (*response.ResponseDepartmentDelete, error) {

	result := &response.ResponseDepartmentDelete{}

	_, err := comp.HttpGet("cgi-bin/department/delete", &object.StringMap{"id": strconv.Itoa(id)}, nil, result)

	return result, err
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/90208
func (comp *Client) List() (*response.ResponseDepartmentList, error) {

	result := &response.ResponseDepartmentList{}

	_, err := comp.HttpGet("cgi-bin/department/list", nil, nil, result)

	return result, err
}
