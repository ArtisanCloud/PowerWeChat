package menu

import (
	"fmt"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/work/menu/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/90232
func (comp *Client) Get() (*response.ResponseMenuGet, error) {

	result := &response.ResponseMenuGet{}

	agentID := (*comp.App).GetConfig().GetInt("agent_id", 0)
	_, err := comp.HttpGet("cgi-bin/menu/get", &object.StringMap{
		"agentid": fmt.Sprintf("%d", agentID),
	}, nil, result)

	return result, err
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/90231
func (comp *Client) Create(data *power.HashMap) (*response.ResponseMenuCreate, error) {

	result := &response.ResponseMenuCreate{}

	agentID := (*comp.App).GetConfig().GetInt("agent_id", 0)
	_, err := comp.HttpPostJson("cgi-bin/menu/create", data.ToHashMap(), &object.StringMap{
		"agentid": fmt.Sprintf("%d", agentID),
	}, nil, result)

	return result, err
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/90233
func (comp *Client) Delete() (*response.ResponseMenuDelete, error) {

	result := &response.ResponseMenuDelete{}

	agentID := (*comp.App).GetConfig().GetInt("agent_id", 0)
	_, err := comp.HttpGet("cgi-bin/menu/delete", &object.StringMap{
		"agentid": fmt.Sprintf("%d", agentID),
	}, nil, result)

	return result, err
}
