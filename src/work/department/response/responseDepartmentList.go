package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/models"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseDepartmentList struct {
	*response.ResponseWork
	Departments []*models.Department `json:"department"`
}
