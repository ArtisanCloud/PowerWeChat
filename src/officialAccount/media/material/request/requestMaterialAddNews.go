package request

import "github.com/ArtisanCloud/power-wechat/src/kernel/power"

type RequestMaterialAddNews struct {
	Articles []*power.HashMap `json:"articles"`
}
