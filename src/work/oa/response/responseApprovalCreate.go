package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseApprovalCreate struct {
	*response.ResponseWork
	SpNo string `json:"sp_no"`
}
