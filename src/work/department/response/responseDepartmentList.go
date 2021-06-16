package response

import (
	"github.com/ArtisanCloud/go-wechat/src/kernel/models"
	"github.com/ArtisanCloud/go-wechat/src/kernel/response"
)

type ResponseDepartmentList struct {
	*response.ResponseWX
	Departments []*models.Department `json:"department"`
}

