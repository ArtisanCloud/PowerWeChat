package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseDepartmentCreate struct {
	*response.ResponseWork
	ID int `json:"id"`
}
