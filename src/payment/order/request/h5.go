package request

type H5Amount struct {
	Total    int    `json:"total"`    // 总金额
	Currency string `json:"currency"` // 货币类型
}

type H5GoodsDetail struct {
	MerchantGoodsId  string `json:"merchant_goods_id"`  // 商户侧商品编码
	WechatPayGoodsId string `json:"wechatpay_goods_id"` // 微信侧商品编码
	GoodsName        string `json:"goods_name"`         // 商品名称
	Quantity         int    `json:"quantity"`           // 商品数量
	UnitPrice        int    `json:"unit_price"`         // 商品单价
}

type H5Detail struct {
	CostPrice   int             `json:"cost_price"`   // 订单原价
	InvoiceId   string          `json:"invoice_id"`   // 商品小票ID
	GoodsDetail []H5GoodsDetail `json:"goods_detail"` // + 单品列表
}

type H5StoreInfo struct {
	Id       string `json:"id"`        // 门店编号
	Name     string `json:"name"`      // 门店名称
	AreaCode string `json:"area_code"` // 地区编码
	Address  string `json:"address"`   // 详细地址
}

type H5H5Info struct {
	Type        string `json:"type"`         // 场景类型
	AppName     string `json:"app_name"`     // 应用名称
	AppUrl      string `json:"app_url"`      // 网站URL
	BundleId    string `json:"bundle_id"`    // IOS平台BundleID
	PackageName string `json:"package_name"` // Android平台PackageName
}

type H5SceneInfo struct {
	PayerClientIp string      `json:"payer_client_ip"` // 用户终端IP
	DeviceId      string      `json:"device_id"`       // 商户端设备号
	StoreInfo     H5StoreInfo `json:"store_info"`      // + 商户门店信息
	H5Info        H5H5Info    `json:"h5_info"`         // + H5场景信息
}

type H5SettleInfo struct {
	ProfitSharing bool `json:"profit_sharing"` // 是否指定分账
}

type RequestH5Prepay struct {
	PrepayBase
	AppID       string       `json:"appid"`        // 应用ID
	Description string       `json:"description"`  // 商品描述
	OutTradeNo  string       `json:"out_trade_no"` // 商户订单号
	TimeExpire  string       `json:"time_expire"`  // 交易结束时间
	Attach      string       `json:"attach"`       // 附加数据
	NotifyUrl   string       `json:"notify_url"`   // 通知地址
	GoodsTag    string       `json:"goods_tag"`    // 订单优惠标记
	Amount      H5Amount     `json:"amount"`       // 订单金额
	Detail      H5Detail     `json:"detail"`       // 优惠功能
	SceneInfo   H5SceneInfo  `json:"scene_info"`   // 场景信息
	SettleInfo  H5SettleInfo `json:"settle_info"`  // 结算信息
}