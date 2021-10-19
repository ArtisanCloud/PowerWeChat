package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseGetAPIDomainIP struct {
	IPList []string `json:"ip_list"`
	*response.ResponseWork
}
