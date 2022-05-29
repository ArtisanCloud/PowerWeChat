package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseUserExportJobs struct {
	*response.ResponseWork

	JobID string `json:"jobid"`
}
