package riskControl

import (
	"context"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/riskControl/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/riskControl/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// 根据提交的用户信息数据获取用户的安全等级 risk_rank，无需用户授权。
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/safety-control-capability/riskControl.getUserRiskRank.html
func (comp *Client) GetUserRiskRank(ctx context.Context, data *request.RequestRiskControl) (*response.ResponseRiskControlGetUserRiskRank, error) {

	result := &response.ResponseRiskControlGetUserRiskRank{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/getuserriskrank", data, nil, nil, result)

	return result, err
}
