package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type ResponseBind struct {
	response.ResponseOpenPlatform

	Userstr string `json:"userstr"`
}

type Member struct {
	UserStr string `json:"userstr"`
}

type ResponseList struct {
	response.ResponseOpenPlatform

	Members []*Member `json:"members"`
}
