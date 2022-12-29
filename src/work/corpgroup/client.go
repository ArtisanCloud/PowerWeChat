package corpgroup

import (
	"context"
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/corpgroup/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// 获取应用共享信息
// https://developer.work.weixin.qq.com/document/path/93403
func (comp *Client) GetAppShareInfo(ctx context.Context, agentID int) (*response.ResponseCorpGroupListAPPShareInfo, error) {

	result := &response.ResponseCorpGroupListAPPShareInfo{}

	params := &object.StringMap{
		"agentid": fmt.Sprintf("%d", agentID),
	}
	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/corpgroup/corp/list_app_share_info", nil, params, nil, result)

	return result, err
}

// 获取下级企业的access_token
// https://developer.work.weixin.qq.com/document/path/93359
func (comp *Client) GetToken(ctx context.Context, corpID string, agentID string) (*response.ResponseCorpGroupGetToken, error) {

	result := &response.ResponseCorpGroupGetToken{}

	params := &object.HashMap{
		"corpid":  corpID,
		"agentid": agentID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/corpgroup/corp/gettoken", params, nil, nil, result)

	return result, err
}

// 获取下级企业的小程序session
// https://developer.work.weixin.qq.com/document/path/93355
func (comp *Client) GetMiniProgramTransferSession(ctx context.Context, userID string, sessionKey string) (*response.ResponseCorpGroupTransferSession, error) {

	result := &response.ResponseCorpGroupTransferSession{}

	params := &object.HashMap{
		"userid":      userID,
		"session_key": sessionKey,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/corpgroup/miniprogram/transfer_session?", params, nil, nil, result)

	return result, err
}
