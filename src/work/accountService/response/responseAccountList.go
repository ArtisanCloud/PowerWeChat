package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseAccountList struct {
	response.ResponseWork

	AccountList []*power.HashMap `json:"account_list"`
}
