package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseDepartmentCreate struct {
	*response.ResponseWork
	ID int `json:"id"`
}
