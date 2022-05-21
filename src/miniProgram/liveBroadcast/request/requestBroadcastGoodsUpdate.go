package request

type RequestBroadcastGoodsUpdate struct {
	GoodsInfo *RequestBroadcastGoodsUpdateInfo `json:"goodsInfo"`
}

type RequestBroadcastGoodsUpdateInfo struct {
	GoodsID         int     `json:"goodsId"`
	CoverImgUrl     string  `json:"coverImgUrl,omitempty"`
	Name            string  `json:"name,omitempty"`
	PriceType       int     `json:"priceType,omitempty'"`
	Price           float64 `json:"price,omitempty"`
	Price2          float64 `json:"price2,omitempty"`
	Url             string  `json:"url,omitempty"`
	ThirdPartyAppid string  `json:"thirdPartyAppid,omitempty"` // 当商品为第三方小程序的商品则填写为对应第三方小程序的appid，自身小程序商品则为''
}
