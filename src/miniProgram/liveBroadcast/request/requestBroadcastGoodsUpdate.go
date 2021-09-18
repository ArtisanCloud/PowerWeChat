package request

import "github.com/ArtisanCloud/power-wechat/src/kernel/power"

type RequestBroadcastGoodsUpdate struct {
	GoodsInfo *power.HashMap `json:"goodsInfo"`
}
