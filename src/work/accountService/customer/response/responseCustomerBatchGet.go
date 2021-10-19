package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseCustomerBatchGet struct {
	*response.ResponseWork

	CustomerList          []*power.HashMap `json:"customer_list"`
	InvalidExternalUserID []*string        `json:"invalid_external_userid"`
}
