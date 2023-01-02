package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/models"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseDepartmentGet struct {
	response.ResponseWork
	Department *models.Department `json:"department"`
}
