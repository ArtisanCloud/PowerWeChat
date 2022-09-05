package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/models"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseDepartmentGet struct {
	*response.ResponseWork
	Department *models.Department `json:"department"`
}
