package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type ResponseBroadcastGetSharedCode struct {
	response.ResponseMiniProgram

	CdnURL    string `json:"cdnUrl"`
	PagePath  string `json:"pagePath"`
	PosterURL string `json:"posterUrl"`
}
