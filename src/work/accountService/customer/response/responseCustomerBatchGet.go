package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseCustomerBatchGet struct {
	*response.ResponseWork

	CustomerList          []*power.HashMap `json:"customer_list"`
	InvalidExternalUserID []*string        `json:"invalid_external_userid"`
}
