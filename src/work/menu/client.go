package menu

import (
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/work/menu/request"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/work/menu/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(&app, nil)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient,
	}, nil
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
func (comp *Client) Create(data *request.RequestMenuSet) (*response.ResponseMenuCreate, error) {

	result := &response.ResponseMenuCreate{}

	agentID := (*comp.App).GetConfig().GetInt("agent_id", 0)
	_, err := comp.HttpPostJson("cgi-bin/menu/create", data, &object.StringMap{
		"agentid": fmt.Sprintf("%d", agentID),
	}, nil, result)

	return result, err
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/90233
func (comp *Client) Delete(agentID int) (*response.ResponseMenuDelete, error) {

	result := &response.ResponseMenuDelete{}

	_, err := comp.HttpGet("cgi-bin/menu/delete", &object.StringMap{
		"agentid": fmt.Sprintf("%d", agentID),
	}, nil, result)

	return result, err
}
