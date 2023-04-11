package request

type CombineSceneInfo struct {
	DeviceID      string `json:"device_id"`       // 商户端设备号
	PayerClientIP string `json:"payer_client_ip"` // 用户终端IP
}

type CombineAmount struct {
	TotalAmount int    `json:"total_amount"` // 标价金额
	Currency    string `json:"currency"`     // 标价币种
}

type CombineSettleInfo struct {
	ProfitSharing bool  `json:"profit_sharing"` // 是否指定分账
	SubsidyAmount int64 `json:"subsidy_amount"` // 补差金额
}

type CombineSubOrders struct {
	MchID       string             `json:"mchid"`                 // 子单商户号
	Attach      string             `json:"attach"`                // 附加数据
	Amount      *CombineAmount     `json:"amount,omitempty"`      // +订单金额
	OutTradeNo  string             `json:"out_trade_no"`          // 子单商户订单号
	GoodsTag    string             `json:"goods_tag"`             // 订单优惠标记
	Description string             `json:"description"`           // 商品描述
	SettleInfo  *CombineSettleInfo `json:"settle_info,omitempty"` // +结算信息
}

type CombineCombinePayerInfo struct {
	OpenID string `json:"openid"` // 用户标识
}

type RequestCombinePrepay struct {
	PrepayBase
	CombineAppID      string                   `json:"combine_appid"`                // 合单商户Appid
	CombineMchID      string                   `json:"combine_mchid"`                // 合单商户号
	CombineOutTradeNo string                   `json:"combine_out_trade_no"`         // 合单商户订单号
	SceneInfo         *CombineSceneInfo        `json:"scene_info,omitempty"`         // 场景信息
	SubOrders         []*CombineSubOrders      `json:"sub_orders,omitempty"`         // 子单信息
	CombinePayerInfo  *CombineCombinePayerInfo `json:"combine_payer_info,omitempty"` // 支付者
	TimeStart         string                   `json:"time_start"`                   // 交易起始时间
	TimeExpire        string                   `json:"time_expire"`                  // 交易结束时间
	NotifyUrl         string                   `json:"notify_url"`                   // 通知地址
}
