package request

import "github.com/ArtisanCloud/powerwechat/src/kernel/power"

type RequestMaterialAddNews struct {
	Articles []*power.HashMap `json:"articles"`
}
