package request

import "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"

type RequestMaterialAddNews struct {
	Articles []*power.HashMap `json:"articles"`
}
