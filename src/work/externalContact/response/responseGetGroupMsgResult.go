package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type ResponseGetGroupMesResult struct {
	response.ResponseWork

	CheckStatus int             `json:"check_status"`
	DetailList  []*ResultDetail `json:"detail_list"`
}

type ResultDetail struct {
	ExternalUserID string `json:"external_userid"` //: "wm_ViZBwAApoZUCOn3JeqlfW1YUme5pg",
	Status         int    `json:"status"`          //: 0,
	Userid         string `json:"userid"`          //: "WangChaoYi"
}
