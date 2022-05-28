package request

type JSAPIAmount struct {
	Total    int    `json:"total"`    // 总金额
	Currency string `json:"currency"` // 货币类型
}

type JSAPIPayer struct {
	OpenID string `json:"openid"` // 用户标识
}

type JSAPIGoodsDetail struct {
	MerchantGoodsID  string `json:"merchant_goods_id"`  // 商户侧商品编码
	WechatPayGoodsID string `json:"wechatpay_goods_id"` // 微信侧商品编码
	GoodsName        string `json:"goods_name"`         // 商品名称
	Quantity         int    `json:"quantity"`           // 商品数量
	UnitPrice        int    `json:"unit_price"`         // 商品单价
}

type JSAPIDetail struct {
	CostPrice   int                 `json:"cost_price"`   // 订单原价
	InvoiceID   string              `json:"invoice_id"`   // 商品小票ID
	GoodsDetail []*JSAPIGoodsDetail `json:"goods_detail"` // + 单品列表
}

type JSAPIStoreInfo struct {
	ID       string `json:"id"`        // 门店编号
	Name     string `json:"name"`      // 门店名称
	AreaCode string `json:"area_code"` // 地区编码
	Address  string `json:"address"`   // 详细地址
}

type JSAPISceneInfo struct {
	PayerClientIP string          `json:"payer_client_ip"` // 用户终端IP
	DeviceID      string          `json:"device_id"`       // 商户端设备号
	StoreInfo     *JSAPIStoreInfo `json:"store_info"`      // + 商户门店信息
}

type JSAPISettleInfo struct {
	ProfitSharing bool `json:"profit_sharing"` // 是否指定分账
}

type RequestJSAPIPrepay struct {
	PrepayBase
	Description string           `json:"description"`           // 商品描述
	OutTradeNo  string           `json:"out_trade_no"`          // 商户订单号
	TimeExpire  string           `json:"time_expire,omitempty"` // 交易结束时间
	Attach      string           `json:"attach,omitempty"`      // 附加数据
	GoodsTag    string           `json:"goods_tag,omitempty"`   // 订单优惠标记
	Amount      *JSAPIAmount     `json:"amount"`                // 订单金额
	Payer       *JSAPIPayer      `json:"payer,omitempty"`       // 支付者
	Detail      *JSAPIDetail     `json:"detail,omitempty"`      // 优惠功能
	SceneInfo   *JSAPISceneInfo  `json:"scene_info,omitempty"`  // 场景信息
	SettleInfo  *JSAPISettleInfo `json:"settle_info,omitempty"` // 结算信息
}
