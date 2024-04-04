package provider

import (
	"context"

	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
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

// 明文corpid转换为加密corpid
// https://developer.work.weixin.qq.com/document/path/95604
func (clt *Client) CorpIDToOpenCorpID(ctx context.Context, corpID string) (string, error) {
	var result struct {
		response.ResponseWork
		OpenCorpID string `json:"open_corpid,omitempty"`
	}
	req := object.HashMap{
		"corpid": corpID,
	}
	_, err := clt.BaseClient.HttpPostJson(ctx, "cgi-bin/service/corpid_to_opencorpid", &req, nil, nil, &result)

	return result.OpenCorpID, err
}
