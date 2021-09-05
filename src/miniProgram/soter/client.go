package soter

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/miniProgram/soter/response"
)

type Client struct {
	*kernel.BaseClient
}

// SOTER 生物认证秘钥签名验证
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/soter/soter.verifySignature.html
func (comp *Client) VerifySignature(openID string, jsonString string, jsonSignature string) (*response.ResponseSoterVerifySignature, error) {

	result := &response.ResponseSoterVerifySignature{}

	data := &object.HashMap{
		"openid":         openID,
		"json_string":    jsonString,
		"json_signature": jsonSignature,
	}

	_, err := comp.HttpPostJson("cgi-bin/soter/verify_signature", data, nil, nil, result)

	return result, err
}
