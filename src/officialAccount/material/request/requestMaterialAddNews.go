package request

import "github.com/ArtisanCloud/PowerWeChat/src/kernel/power"

type RequestMaterialAddNews struct {
	Articles []*power.HashMap `json:"articles"`
}
