package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseAccountList struct {
	*response.ResponseWork

	AccountList []*power.HashMap `json:"account_list"`
}
