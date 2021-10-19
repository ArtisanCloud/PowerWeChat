package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseDepartmentCreate struct {
	*response.ResponseWork
	ID int `json:"id"`
}
