package workbench

import (
	"fmt"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/work/agent/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/90205
func (comp *Client) SetWorkbenchTemplate(data *power.HashMap) (*response.ResponseAgentSetWorkbenchTemplate, error) {

	result := &response.ResponseAgentSetWorkbenchTemplate{}

	_, err := comp.HttpPostJson("cgi-bin/agent/set_workbench_template", data, nil, nil, result)

	return result, err
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/90206
func (comp Client) GetWorkbenchTemplate(agentID int) (*response.ResponseAgentGetWorkbenchTemplate, error) {

	result := &response.ResponseAgentGetWorkbenchTemplate{}

	_, err := comp.HttpGet("cgi-bin/agent/get_workbench_template", &object.StringMap{
		"agentid": fmt.Sprintf("%d", agentID),
	}, nil, result)

	return result, err
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/90207
func (comp Client) SetWorkbenchData(data *power.HashMap) (*response.ResponseAgentSetWorkbenchData, error) {

	result := &response.ResponseAgentSetWorkbenchData{}

	_, err := comp.HttpGet("cgi-bin/agent/set_workbench_data", data, nil, result)

	return result, err
}
