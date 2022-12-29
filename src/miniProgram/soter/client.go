package soter

import (
	"context"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/soter/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/soter/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// SOTER 生物认证秘钥签名验证
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/soter/soter.verifySignature.html
func (comp *Client) VerifySignature(ctx context.Context, options *request.RequestSoter) (*response.ResponseSoterVerifySignature, error) {

	result := &response.ResponseSoterVerifySignature{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/soter/verify_signature", options, nil, nil, result)

	return result, err
}
