package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseApprovalGetData struct {
	*response.ResponseWork
	Count     int            `json:"count"`
	Total     int            `json:"total"`
	NextSPNum int            `json:"next_spnum"`
	Data      *power.HashMap `json:"data"`
}
