package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseDepartmentCreate struct {
	*response.ResponseWork
	ID int `json:"id"`
}
