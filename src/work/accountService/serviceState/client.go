package serviceState

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/src/work/accountService/serviceState/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

// 获取会话状态
// https://work.weixin.qq.com/api/doc/90000/90135/94669
func (comp *Client) Get(openKFID string, externalUserID string) (*response.ResponseServiceStateGet, error) {

	result := &response.ResponseServiceStateGet{}

	options := &object.HashMap{
		"open_kfid":       openKFID,
		"external_userid": externalUserID,
	}

	_, err := comp.HttpPostJson("cgi-bin/kf/service_state/get", options, nil, nil, result)

	return result, err
}


// 变更会话状态
// https://work.weixin.qq.com/api/doc/90000/90135/94669
func (comp *Client) Trans(openKFID string, externalUserID string, serviceState int, servicerUserID string ) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	options := &object.HashMap{
		"open_kfid":       openKFID,
		"external_userid": externalUserID,
		"service_state": serviceState,
		"servicer_userid": servicerUserID,
	}

	_, err := comp.HttpPostJson("cgi-bin/kf/service_state/trans", options, nil, nil, result)

	return result, err
}