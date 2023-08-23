package request

type VirtualPaymentOrderRequest struct {
	SessionKey string `json:"session_key"`  // 文档已更新，现在允许存session key 过期时间官方写的是三天，配合前端checkSessionKey组件经测试 官方有bug
	ProductId  int64  `json:"product_id"`   // 商品id
	Price      int64  `json:"price"`        // 金额 单位分
	OutTradeNo string `json:"out_trade_no"` // 订单号
	Attach     string `json:"attach"`       // 附加信息
}

type UploadProductsRequest struct {
	Env        int8        `json:"env"` // 0正式 1测试
	UploadItem []*GoodItem `json:"upload_item"`
}

type GoodItem struct {
	Id      string `json:"id"`       // 商品id 同product id 长度(0,64]，字符只允许使用字母、数字、'_'、'-'
	Name    string `json:"name"`     // 商品名称 长度(0,1024] 不能有特殊字符
	Price   int64  `json:"price"`    // 道具单价，单位分，需要大于0
	Remake  string `json:"remake"`   // 道具备注，长度(0,1024]
	ItemUrl string `json:"item_url"` // 道具图片的url地址，当前仅支持jpg,png等格式 使用查询上传结果接口会发现，经常会报错上传图片失败 需要多传几次
}
