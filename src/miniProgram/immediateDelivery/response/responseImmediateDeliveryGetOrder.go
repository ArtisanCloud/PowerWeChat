package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseImmediateDeliveryGetOrder struct {
	response.ResponseMiniProgram

	OrderStatus int     `json:"order_status"` //	配送状态，枚举值
	WaybillID   string  `json:"waybill_id"`   //	配送单号
	RiderName   string  `json:"rider_name"`   //	骑手姓名
	RiderPhone  string  `json:"rider_phone"`  //	骑手电话
	RiderLng    float64 `json:"rider_lng"`    //	骑手位置经度, 配送中时返回
	RiderLat    float64 `json:"rider_lat"`    //	骑手位置纬度, 配送中时返回
	ReachTime   int64   `json:"reach_time"`   //	预计还剩多久送达时间, 配送中时返回，单位秒， 已取货配送中需返回，比如5分钟后送达，填300

}
