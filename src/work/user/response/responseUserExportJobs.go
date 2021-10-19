package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseUserExportJobs struct {
	*response.ResponseWork

	JobID string `json:"jobid"`
}
