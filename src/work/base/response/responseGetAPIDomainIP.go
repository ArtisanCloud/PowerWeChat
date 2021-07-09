package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseGetAPIDomainIP struct {
	IPList []string `json:"ip_list"`
	*response.ResponseWX
}

