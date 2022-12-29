package agent

import (
	"context"
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/agent/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/agent/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
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

// https://developer.work.weixin.qq.com/document/path/90227
func (comp *Client) Get(ctx context.Context, agentID int) (*response.ResponseAgentGet, error) {

	result := &response.ResponseAgentGet{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/agent/get", nil, &object.StringMap{
		"agentid": fmt.Sprintf("%d", agentID),
	}, nil, result)

	return result, err
}

// https://developer.work.weixin.qq.com/document/path/90228
func (comp *Client) Set(ctx context.Context, data *request.RequestAgentSet) (*response.ResponseAgentSet, error) {

	result := &response.ResponseAgentSet{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/agent/set", data, nil, nil, result)

	return result, err
}

// https://developer.work.weixin.qq.com/document/path/90227
func (comp *Client) List(ctx context.Context) (*response.ResponseAgentList, error) {

	result := &response.ResponseAgentList{}

	_, err := comp.BaseClient.HttpGet(ctx, "cgi-bin/agent/list", nil, nil, result)

	return result, err
}
