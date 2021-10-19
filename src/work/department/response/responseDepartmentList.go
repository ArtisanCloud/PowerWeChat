package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/models"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseDepartmentList struct {
	*response.ResponseWork
	Departments []*models.Department `json:"department"`
}
