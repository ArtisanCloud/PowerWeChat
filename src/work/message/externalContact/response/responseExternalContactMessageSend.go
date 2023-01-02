package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseExternalContactMessageSend struct {
	response.ResponseWork

	InvalidExternalUser  []string `json:"invalid_external_user"`
	InvalidParentUserID  []string `json:"invalid_parent_userid"`
	InvalidStudentUserID []string `json:"invalid_student_userid"`
	InvalidParty         []string `json:"invalid_party"`
}
