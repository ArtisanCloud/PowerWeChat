package base

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/base/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/base/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/user-info/auth.getPaidUnionId.html
func (comp *Client) GetPaidUnionID(ctx context.Context, data *request.RequestGetPaidUnionID) (*response.ResponseAuthGetPaidUnionID, error) {

	result := &response.ResponseAuthGetPaidUnionID{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}
	_, err = comp.BaseClient.HttpGet(ctx, "wxa/getpaidunionid", params, nil, result)

	return result, err
}

// 检查加密信息是否由微信生成（当前只支持手机号加密数据），只能检测最近3天生成的加密数据
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/user-info/auth.checkEncryptedData.html
func (comp *Client) CheckEncryptedData(ctx context.Context, encryptedMsgHash string) (*response.ResponseAuthCheckEncryptedData, error) {

	result := &response.ResponseAuthCheckEncryptedData{}

	options := &object.HashMap{
		"encrypted_msg_hash": encryptedMsgHash,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/business/checkencryptedmsg", options, nil, nil, result)

	return result, err
}
