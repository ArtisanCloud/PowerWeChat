package menu

import (
	"context"
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/menu/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/menu/response"
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

// https://developer.work.weixin.qq.com/document/path/90232
func (comp *Client) Get(ctx context.Context) (*response.ResponseMenuGet, error) {

	result := &response.ResponseMenuGet{}

	agentID := (*comp.BaseClient.App).GetConfig().GetInt("agent_id", 0)
	_, err := comp.BaseClient.HttpGet(ctx, "cgi-bin/menu/get", &object.StringMap{
		"agentid": fmt.Sprintf("%d", agentID),
	}, nil, result)

	return result, err
}

// https://developer.work.weixin.qq.com/document/path/90231
func (comp *Client) Create(ctx context.Context, data *request.RequestMenuSet) (*response.ResponseMenuCreate, error) {

	result := &response.ResponseMenuCreate{}

	agentID := (*comp.BaseClient.App).GetConfig().GetInt("agent_id", 0)
	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/menu/create", data, &object.StringMap{
		"agentid": fmt.Sprintf("%d", agentID),
	}, nil, result)

	return result, err
}

// https://developer.work.weixin.qq.com/document/path/90233
func (comp *Client) Delete(ctx context.Context, agentID int) (*response.ResponseMenuDelete, error) {

	result := &response.ResponseMenuDelete{}

	_, err := comp.BaseClient.HttpGet(ctx, "cgi-bin/menu/delete", &object.StringMap{
		"agentid": fmt.Sprintf("%d", agentID),
	}, nil, result)

	return result, err
}
