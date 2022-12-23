package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type DataMaterialUpload struct {
	PicURL string `json:"pic_url"`
}

type ResponseMaterialUpload struct {
	*response.ResponseOfficialAccount

	Data DataMaterialUpload `json:"data"`
}
