package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseUserExportJobs struct {
	*response.ResponseWork

	JobID string `json:"jobid"`
}
