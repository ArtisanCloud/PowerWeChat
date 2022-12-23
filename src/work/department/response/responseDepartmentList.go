package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/models"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseDepartmentList struct {
	response.ResponseWork
	Departments []*models.Department `json:"department"`
}

type DepartmentID struct {
	ID       int `json:"id"`
	ParentID int `json:"parentid"`
	Order    int `json:"order"`
}

type ResponseDepartmentIDList struct {
	response.ResponseWork

	DepartmentIDs []DepartmentID `json:"department_id"`
}
