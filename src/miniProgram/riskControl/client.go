package riskControl

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel"
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/miniProgram/riskControl/response"
)

type Client struct {
	*kernel.BaseClient
}

// 根据提交的用户信息数据获取用户的安全等级 risk_rank，无需用户授权。
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/safety-control-capability/riskControl.getUserRiskRank.html
func (comp *Client) GetUserRiskRank(data *power.HashMap) (*response.ResponseRiskControlGetUserRiskRank, error) {

	result := &response.ResponseRiskControlGetUserRiskRank{}

	_, err := comp.HttpPostJson("wxa/getuserriskrank", data, nil, nil, result)

	return result, err
}
