package base

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/miniProgram/base/response"
)

type Client struct {
	*kernel.BaseClient
}

// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/user-info/auth.getPaidUnionId.html
func (comp *Client) getPaidUnionID(openID string, option *object.StringMap) (*response.ResponseGetPaidUnionID, error) {

	result := &response.ResponseGetPaidUnionID{}

	params := &object.StringMap{
		"openid": openID,
	}

	params = object.MergeStringMap(params, option)

	_, err := comp.HttpGet("wxa/getpaidunionid", params, nil, result)

	return result, err
}

// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/user-info/auth.checkEncryptedData.html
func (comp *Client) CheckEncryptedMsg(encryptedMsgHash string) (*response.ResponseCheckEncryptedMsg, error) {

	result := &response.ResponseCheckEncryptedMsg{}

	options := &object.HashMap{
		"encrypted_msg_hash": encryptedMsgHash,
	}

	_, err := comp.HttpPostJson("wxa/business/checkencryptedmsg", options, nil, nil, result)

	return result, err
}


// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getPerformanceData.html
func (comp *Client) GetPerformanceData(options *power.HashMap) (*response.ResponseGetPerformanceData, error) {

	result := &response.ResponseGetPerformanceData{}

	_, err := comp.HttpPostJson("wxa/business/performance/boot", options, nil, nil, result)

	return result, err
}