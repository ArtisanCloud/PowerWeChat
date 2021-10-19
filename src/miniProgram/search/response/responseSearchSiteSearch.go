package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseSearchSiteSearch struct {
	*response.ResponseMiniProgram

	Items        []*power.HashMap `json:"items"`
	HasNextPage  int              `json:"has_next_page"`
	HitCount     int              `json:"hit_count"`
	NextPageInfo string           `json:"next_page_info"`
}
