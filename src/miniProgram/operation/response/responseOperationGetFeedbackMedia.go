package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseOperationGetFeedbackMedia struct {
	*response.ResponseMiniProgram
	RequestDomain   []string `json:"requestdomain"`
	WSRequestDomain []string `json:"wsrequestdomain"`
	UploadDomain    []string `json:"uploaddomain"`
	DownloadDomain  []string `json:"downloaddomain"`
	UdpDomain       []string `json:"udpdomain"`
	BizDomain       []string `json:"bizdomain"`
}
