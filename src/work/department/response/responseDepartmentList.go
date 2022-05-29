package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/models"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseDepartmentList struct {
	*response.ResponseWork
	Departments []*models.Department `json:"department"`
}
