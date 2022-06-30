package response

import "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"

type Interface struct {
	ApiName   string `json:"api_name"`
	ApiChName string `json:"api_ch_name"`
	ApiDesc   string `json:"api_desc"`
	Status    int    `json:"status"`
	ApiLink   string `json:"api_link"`
	GroupName string `json:"group_name"`
}

type ResponseGet struct {
	response.ResponseOpenPlatform

	InterfaceList []*Interface `json:"interface_list"`
}

type ResponseUpload struct {
	response.ResponseOpenPlatform

	ExtFileMediaID string `json:"ext_file_media_id"`
}
