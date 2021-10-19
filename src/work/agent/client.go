package agent

import (
	"fmt"
	"github.com/ArtisanCloud/powerwechat/src/kernel"
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/work/agent/request"
	"github.com/ArtisanCloud/powerwechat/src/work/agent/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/90227
func (comp *Client) Get(agentID int) (*response.ResponseAgentGet, error) {

	result := &response.ResponseAgentGet{}

	_, err := comp.HttpPostJson("cgi-bin/agent/get", nil, &power.StringMap{
		"agentid": fmt.Sprintf("%d", agentID),
	}, nil, result)

	return result, err
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/90228
func (comp *Client) Set(data *request.RequestAgentSet) (*response.ResponseAgentSet, error) {

	result := &response.ResponseAgentSet{}

	_, err := comp.HttpPostJson("cgi-bin/agent/set", data, nil, nil, result)

	return result, err
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/90227
func (comp *Client) List() (*response.ResponseAgentList, error) {

	result := &response.ResponseAgentList{}

	_, err := comp.HttpGet("cgi-bin/agent/list", nil, nil, result)

	return result, err
}
