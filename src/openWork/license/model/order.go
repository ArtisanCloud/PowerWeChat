package model

type AccountCount struct {
	// BaseCount 基础账号个数，最多1000000个。(若企业为服务商测试企业，最多购买1000个)
	BaseCount int `json:"base_count,omitempty"`
	// ExternalContactCount 互通账号个数，最多1000000个。(若企业为服务商测试企业，最多购买1000个)
	ExternalContactCount int `json:"external_contact_count,omitempty"`
}

type AccountDuration struct {
	// Month 购买的月数，每个月按照31天计算
	Month int `json:"month,omitempty"`
	// Days 购买的天数
	Days int `json:"days,omitempty"`
	// NewExpireTime 下单续期账号中指定新过期时间时返回
	NewExpireTime int64 `json:"new_expire_time,omitempty"`
}

type Order struct {
	// OrderID 订单id
	OrderID string `json:"order_id,omitempty"`
	// OrderType 订单类型
	// 1：购买账号
	// 2：续期账号
	// 5：历史企业迁移订单
	// 8：多企业新购订单(只返回父订单，且仅当corpid不填时返回)
	OrderType int `json:"order_type,omitempty"`
	// OrderStatus 订单状态，0：待支付，1：已支付，2：已取消（未支付，订单已关闭）3：未支付，订单已过期，4：申请退款中，5：退款成功，6：退款被拒绝，7：订单已失效（将企业从服务商测试企业列表中移除时会将对应测试企业的所有测试订单置为已失效）
	OrderStatus int `json:"order_status,omitempty"`
	// CorpID 客户企业id，返回加密的corpid
	CorpID string `json:"corp_id,omitempty"`
	// Price 订单金额，单位分
	Price int64 `json:"price,omitempty"`
	// AccountCount 订单的账号数详情
	AccountCount *AccountCount `json:"account_count,omitempty"`
	// AccountDuration 账号购买时长
	AccountDuration *AccountDuration `json:"account_duration,omitempty"`
	// CreateTime 创建时间
	CreateTime int64 `json:"create_time,omitempty"`
	// PayTime 支付时间。迁移订单不返回该字段
	PayTime int64 `json:"pay_time,omitempty"`
}
