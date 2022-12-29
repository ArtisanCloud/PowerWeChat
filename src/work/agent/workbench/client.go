package workbench

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

// https://developer.work.weixin.qq.com/document/path/90205
func (comp *Client) SetWorkbenchTemplate(ctx context.Context, data *request.RequestSetWorkbenchTemplate) (*response.ResponseAgentSetWorkbenchTemplate, error) {

	result := &response.ResponseAgentSetWorkbenchTemplate{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/agent/set_workbench_template", data, nil, nil, result)

	return result, err
}

// https://developer.work.weixin.qq.com/document/path/90206
func (comp Client) GetWorkbenchTemplate(ctx context.Context, agentID int) (*response.ResponseAgentGetWorkbenchTemplate, error) {

	result := &response.ResponseAgentGetWorkbenchTemplate{}

	_, err := comp.BaseClient.HttpGet(ctx, "cgi-bin/agent/get_workbench_template", &object.StringMap{
		"agentid": fmt.Sprintf("%d", agentID),
	}, nil, result)

	return result, err
}

// https://developer.work.weixin.qq.com/document/path/92535
func (comp Client) SetWorkbenchData(ctx context.Context, data *request.RequestSetWorkBenchData) (*response.ResponseAgentSetWorkbenchData, error) {

	result := &response.ResponseAgentSetWorkbenchData{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}
	_, err = comp.BaseClient.HttpPost(ctx, "cgi-bin/agent/set_workbench_data", params, nil, result)

	return result, err
}
