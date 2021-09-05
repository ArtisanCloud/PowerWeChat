package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseOperationGetDomainInfo struct {
	*response.ResponseMiniProgram
	RequestDomain   []string `json:"requestdomain"`
	WSRequestDomain []string `json:"wsrequestdomain"`
	UploadDomain    []string `json:"uploaddomain"`
	DownloadDomain  []string `json:"downloaddomain"`
	UdpDomain       []string `json:"udpdomain"`
	BizDomain       []string `json:"bizdomain"`
}
