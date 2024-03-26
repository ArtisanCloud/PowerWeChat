package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/license/model"
)

type ResponseListOrder struct {
	response.ResponseWork
	// NextCursor 分页游标，再下次请求时填写以获取之后分页的记录
	NextCursor string `json:"next_cursor,omitempty"`
	// HasMore 是否有更多。 0: 没有， 1: 有
	HasMore int `json:"has_more,omitempty"`
	// OrderList 订单列表
	OrderList []model.Order `json:"order_list,omitempty"`
}

type ResponseGetOrder struct {
	response.ResponseWork
	Order *model.Order `json:"order,omitempty"`
}
