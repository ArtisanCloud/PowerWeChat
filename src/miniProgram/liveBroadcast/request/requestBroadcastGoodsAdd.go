package request

import "github.com/ArtisanCloud/power-wechat/src/kernel/power"

type RequestBroadcastGoodsAdd struct {
	GoodsInfo *power.HashMap `json:"goodsInfo"`
}
