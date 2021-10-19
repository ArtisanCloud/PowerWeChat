package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseLinkCorpGetPermList struct {
	*response.ResponseWork

	UserIDs []string `json:"userids"`
	DepartmentIDs []string `json:"department_ids"`
}
