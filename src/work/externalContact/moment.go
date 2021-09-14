package externalContact

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/request"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/response"
)

type Moment struct {
	*kernel.BaseClient
}

func NewMomentClient(app kernel.ApplicationInterface) *Moment {
	return &Moment{
		kernel.NewBaseClient(&app, nil),
	}
}

// 获取企业全部的发表列表
// https://work.weixin.qq.com/api/doc/90000/90135/93333
func (comp *GroupChat) GetMomentList(params *request.RequestGetMomentList) (*response.ResponseGetMomentList, error) {

	result := &response.ResponseGetMomentList{}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/get_moment_list", params, nil, nil, result)

	return result, err
}

// 客户朋友圈规则组管理
// https://work.weixin.qq.com/api/doc/90000/90135/94890
func (comp *GroupChat) MomentStrategyList(cursor string, limit int) (*response.ResponseMomentStrategyList, error) {

	result := &response.ResponseMomentStrategyList{}

	options := &object.HashMap{
		"cursor": cursor,
		"limit":  limit,
	}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/moment_strategy/list", options, nil, nil, result)

	return result, err
}
