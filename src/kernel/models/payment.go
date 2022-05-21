package models

import "time"

// --- Transaction models ----

const WX_TRADE_STATE_ = "JSAPI"            //公众号支付
const WX_TRADE_STATE_NATIVE = "NATIVE"     //扫码支付
const WX_TRADE_STATE_APP = "APP"           //APP支付
const WX_TRADE_STATE_MICROPAY = "MICROPAY" //付款码支付
const WX_TRADE_STATE_MWEB = "MWEB"         //H5支付
const WX_TRADE_STATE_FACEPAY = "FACEPAY"   //刷脸支付

const WX_TRADE_STATE_SUCCESS = "SUCCESS"       // 支付成功
const WX_TRADE_STATE_REFUND = "REFUND"         // 转入退款
const WX_TRADE_STATE_NOTPAY = "NOTPAY"         // 未支付
const WX_TRADE_STATE_CLOSED = "CLOSED"         // 已关闭
const WX_TRADE_STATE_REVOKED = "REVOKED"       // 已撤销（付款码支付）
const WX_TRADE_STATE_USERPAYING = "USERPAYING" // 用户支付中（付款码支付）
const WX_TRADE_STATE_PAYERROR = "PAYERROR"     // 支付失败(其他原因，如银行返回失败)

// Transaction
type Transaction struct {
	Amount          *TransactionAmount `json:"amount,omitempty"`
	AppID           string             `json:"appid,omitempty"`
	Attach          string             `json:"attach,omitempty"`
	BankType        string             `json:"bank_type,omitempty"`
	MchID           string             `json:"mchid,omitempty"`
	OutTradeNo      string             `json:"out_trade_no,omitempty"`
	Payer           *TransactionPayer  `json:"payer,omitempty"`
	PromotionDetail []PromotionDetail  `json:"promotion_detail,omitempty"`
	SuccessTime     string             `json:"success_time,omitempty"`
	TradeState      string             `json:"trade_state,omitempty"`
	TradeStateDesc  string             `json:"trade_state_desc,omitempty"`
	TradeType       string             `json:"trade_type,omitempty"`
	TransactionID   string             `json:"transaction_id,omitempty"`
}

// TransactionAmount
type TransactionAmount struct {
	Currency      string `json:"currency,omitempty"`
	PayerCurrency string `json:"payer_currency,omitempty"`
	PayerTotal    int64  `json:"payer_total,omitempty"`
	Total         int64  `json:"total,omitempty"`
}

// TransactionPayer
type TransactionPayer struct {
	OpenID string `json:"openid,omitempty"`
}

// PromotionDetail
type PromotionDetail struct {
	// 券ID
	CouponID string `json:"coupon_id,omitempty"`
	// 优惠名称
	Name string `json:"name,omitempty"`
	// GLOBAL：全场代金券；SINGLE：单品优惠
	Scope string `json:"scope,omitempty"`
	// CASH：充值；NOCASH：预充值。
	Type string `json:"type,omitempty"`
	// 优惠券面额
	Amount int64 `json:"amount,omitempty"`
	// 活动ID，批次ID
	StockID string `json:"stock_id,omitempty"`
	// 单位为分
	WechatpayContribute int64 `json:"wechatpay_contribute,omitempty"`
	// 单位为分
	MerchantContribute int64 `json:"merchant_contribute,omitempty"`
	// 单位为分
	OtherContribute int64 `json:"other_contribute,omitempty"`
	// CNY：人民币，境内商户号仅支持人民币。
	Currency    string                 `json:"currency,omitempty"`
	GoodsDetail []PromotionGoodsDetail `json:"goods_detail,omitempty"`
}

// PromotionGoodsDetail
type PromotionGoodsDetail struct {
	// 商品编码
	GoodsID string `json:"goods_id"`
	// 商品数量
	Quantity int64 `json:"quantity"`
	// 商品价格
	UnitPrice int64 `json:"unit_price"`
	// 商品优惠金额
	DiscountAmount int64 `json:"discount_amount"`
	// 商品备注
	GoodsRemark string `json:"goods_remark,omitempty"`
}

// --- Refund models ----
// Refund
type Refund struct {
	// 原支付交易对应的商户订单号
	MchID string `json:"mchid"`
	// 微信支付交易订单号
	TransactionID string `json:"transaction_id"`
	// 原支付交易对应的商户订单号
	OutTradeNo string `json:"out_trade_no"`
	// 微信支付退款号
	RefundID string `json:"refund_id"`
	// 商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。
	OutRefundNo string `json:"out_refund_no"`
	// 退款状态，枚举值
	RefundStatus string `json:"refund_status"`
	// 退款成功时间，退款状态status为SUCCESS（退款成功）时，返回该字段。遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日13点29分35秒。
	SuccessTime *time.Time `json:"success_time,omitempty"`
	// 取当前退款单的退款入账方，有以下几种情况： 1）退回银行卡：{银行名称}{卡类型}{卡尾号} 2）退回支付用户零钱:支付用户零钱 3）退还商户:商户基本账户商户结算银行账户 4）退回支付用户零钱通:支付用户零钱通
	UserReceivedAccount string `json:"user_received_account"`
	// 金额详细信息
	Amount *Amount `json:"amount"`
}

// Amount
type Amount struct {
	// 订单总金额，单位为分
	Total int64 `json:"total"`
	// 退款标价金额，单位为分，可以做部分退款
	Refund int64 `json:"refund"`
	// 现金支付金额，单位为分，只能为整数
	PayerTotal int64 `json:"payer_total"`
	// 退款给用户的金额，不包含所有优惠券金额
	PayerRefund int64 `json:"payer_refund"`
	// 去掉非充值代金券退款金额后的退款金额，单位为分，退款金额=申请退款金额-非充值代金券退款金额，退款金额<=申请退款金额

}
