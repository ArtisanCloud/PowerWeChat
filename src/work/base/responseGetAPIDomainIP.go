package base

import (
	"github.com/ArtisanCloud/go-wechat/src/kernel/response"
)

type ResponseGetAPIDomainIP struct {
	IPList []string `json:"ip_list"`
	*response.ResponseWX
}

