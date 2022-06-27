package response

import "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"

type ResponseSession struct {
	response.ResponseWork

	CorpID     string `json:"corpid"`
	UserID     string `json:"userid"`
	SessionKey string `json:"session_key"`
}
