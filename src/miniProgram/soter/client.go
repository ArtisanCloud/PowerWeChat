package soter

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/src/miniProgram/soter/request"
	"github.com/ArtisanCloud/PowerWeChat/src/miniProgram/soter/response"
)

type Client struct {
	*kernel.BaseClient
}

// SOTER 生物认证秘钥签名验证
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/soter/soter.verifySignature.html
func (comp *Client) VerifySignature(options *request.RequestSoter) (*response.ResponseSoterVerifySignature, error) {

	result := &response.ResponseSoterVerifySignature{}

	_, err := comp.HttpPostJson("cgi-bin/soter/verify_signature", options, nil, nil, result)

	return result, err
}
