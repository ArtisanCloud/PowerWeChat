package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type InvalidSenderList struct {
	UserList       []string `json:"user_list"`
	DepartmentList []int    `json:"department_list"`
}

type InvalidExternalContactList struct {
	TagList []string `json:"tag_list"`
}

type Result struct {
	response.ResponseWork
	MomentId                   string                     `json:"moment_id"`
	InvalidSenderList          InvalidSenderList          `json:"invalid_sender_list"`
	InvalidExternalContactList InvalidExternalContactList `json:"invalid_external_contact_list"`
}

type ResponseAddMomentTask struct {
	response.ResponseWork
	JobId int `json:"jobid"`
}
