package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/models"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseDepartmentList struct {
	*response.ResponseWork
	Departments []*models.Department `json:"department"`
}
