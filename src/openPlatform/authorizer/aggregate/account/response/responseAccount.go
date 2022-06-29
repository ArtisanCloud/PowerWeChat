package response

import "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"

type ResponseCreate struct {
	response.ResponseOpenPlatform

	OpenAppID string `json:"open_appid"`
}

type ResponseGetBinding struct {
	response.ResponseOpenPlatform

	OpenAppid string `json:"open_appid"`
}
