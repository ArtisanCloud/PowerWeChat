package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseGetAPIDomainIP struct {
	IPList []string `json:"ip_list"`
	*response.ResponseWork
}
