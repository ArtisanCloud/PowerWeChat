package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseApprovalCreate struct {
	*response.ResponseWork
	SpNo string `json:"sp_no"`
}
