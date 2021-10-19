package response

import "github.com/ArtisanCloud/powerwechat/src/kernel/response"

type ResponseBroadcastGetSharedCode struct {
	*response.ResponseMiniProgram

	CdnURL    string `json:"cdnUrl"`
	PagePath  string `json:"pagePath"`
	PosterURL string `json:"posterUrl"`
}
