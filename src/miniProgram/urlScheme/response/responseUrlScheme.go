package response

import "github.com/ArtisanCloud/PowerWeChat/src/kernel/response"

type ResponseUrlScheme struct {
	*response.ResponseMiniProgram

	OpenLink string `json:"openlink"`
}
