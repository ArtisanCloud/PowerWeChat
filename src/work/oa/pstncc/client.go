package pstncc

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/src/work/oa/pstncc/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

// 发起语音电话
// https://work.weixin.qq.com/api/doc/90000/90135/91627
func (comp *Client) Call(calleeUserID []string) (*response.ResponsePSTNCCCall, error) {

	result := &response.ResponsePSTNCCCall{}

	options := &object.HashMap{
		"callee_userid": calleeUserID,
	}

	_, err := comp.HttpPostJson("cgi-bin/pstncc/call", options, nil, nil, result)

	return result, err
}

// 获取接听状态
// https://work.weixin.qq.com/api/doc/90000/90135/91628
func (comp *Client) GetStates(calleeUserID string, callID string) (*response.ResponsePSTNCCGetStates, error) {

	result := &response.ResponsePSTNCCGetStates{}

	options := &object.HashMap{
		"callee_userid": calleeUserID,
		"callid": callID,
	}

	_, err := comp.HttpPostJson("cgi-bin/pstncc/getstates", options, nil, nil, result)

	return result, err
}
