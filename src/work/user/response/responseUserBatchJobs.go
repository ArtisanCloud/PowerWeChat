package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseUserBatchJobs struct {
	response.ResponseWork

	JobID string `json:"jobid"`
}
