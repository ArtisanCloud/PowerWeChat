package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseCorpGroupTransferSession struct {
	*response.ResponseWork
	Userid     string `json:"userid"`
	SessionKey string `json:"session_key"`
}
