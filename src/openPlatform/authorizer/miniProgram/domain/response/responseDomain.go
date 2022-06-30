package response

import "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"

type ResponseModify struct {
	response.ResponseOpenPlatform

	RequestDomain          []interface{} `json:"requestdomain"`
	WSRequestDomain        []interface{} `json:"wsrequestdomain"`
	UploadDomain           []interface{} `json:"uploaddomain"`
	DownloadDomain         []interface{} `json:"downloaddomain"`
	UDPDomain              []interface{} `json:"udpdomain"`
	TCPDomain              []interface{} `json:"tcpdomain"`
	InvalidRequestDomain   []interface{} `json:"invalid_requestdomain"`
	InvalidWSRequestDomain []string      `json:"invalid_wsrequestdomain"`
	InvalidUploadDomain    []string      `json:"invalid_uploaddomain"`
	InvalidDownloadDomain  []interface{} `json:"invalid_downloaddomain"`
	InvalidUDPDomain       []string      `json:"invalid_udpdomain"`
	InvalidTCPDomain       []string      `json:"invalid_tcpdomain"`
	NoICPDomain            []string      `json:"no_icp_domain"`
}
