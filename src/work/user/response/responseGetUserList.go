package response

import (
	"github.com/ArtisanCloud/PowerSocialite/v3/src/models"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseGetUserList struct {
	response.ResponseWork

	//UserList []*response2.RequestUserDetail `json:"userlist"`
	UserList []*models.Employee `json:"userlist"`
}

type DepartmentUser struct {
	UserID     string `json:"userid"`
	Department int    `json:"department"`
}

type ResponseListID struct {
	response.ResponseWork

	NextCursor string            `json:"next_cursor"`
	DeptUser   []*DepartmentUser `json:"dept_user"`
}
