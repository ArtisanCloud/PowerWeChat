package approval

import (
	"context"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/oa/approval/request"
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

// 更新审批模板
// https://developer.work.weixin.qq.com/document/path/97438
func (comp *Client) UpdateTemplate(ctx context.Context, options *request.RequestUpdateTemplate) (*response.ResponseWork, error) {

	result := &response.ResponseWork{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/oa/approval/update_template", options, nil, nil, result)

	return result, err
}
