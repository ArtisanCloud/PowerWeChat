package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseUserBatchJobs struct {
	*response.ResponseWork

	JobID string `json:"jobid"`
}
