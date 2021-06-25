package response

import "github.com/ArtisanCloud/go-wechat/src/kernel/response"

type ResponseGetList struct {
	response.ResponseWX
	ExternalUserID []string `json:"external_userid"` // ["woAJ2GCAAAXtWyujaWJHDDGi0mACAAA","wmqfasd1e1927831291723123109rAAA"]
}