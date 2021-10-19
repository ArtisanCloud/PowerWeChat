package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseUserBatchJobs struct {
	*response.ResponseWork

	JobID string `json:"jobid"`
}
