package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseOCRIDCard struct {
	*response.ResponseMiniProgram
	Type        string `json:"type"`
	Name        string `json:"name"`
	Id          string `json:"id"`
	Addr        string `json:"addr"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
	ValidDate   string `json:"valid_date"`
}
