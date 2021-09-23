package request

type RequestBroadcastGoodsAdd struct {
	GoodsInfo *RequestBroadcastGoodsAddInfo `json:"goodsInfo"`
}

type RequestBroadcastGoodsAddInfo struct {
	CoverImgUrl     string  `json:"coverImgUrl"`
	Name            string  `json:"name"`
	PriceType       int     `json:"priceType"`
	Price           float64 `json:"price"`
	Price2          float64 `json:"price2"`
	Url             string  `json:"url"`
	ThirdPartyAppid string  `json:"thirdPartyAppid"` // 当商品为第三方小程序的商品则填写为对应第三方小程序的appid，自身小程序商品则为''
}
