package request

type RequestModify struct {
	Action          string   `json:"action"`
	RequestDomain   []string `json:"requestdomain"`
	WSRequestDomain []string `json:"wsrequestdomain"`
	UploadDomain    []string `json:"uploaddomain"`
	DownloadDomain  []string `json:"downloaddomain"`
	UDPDomain       []string `json:"udpdomain"`
	TCPDomain       []string `json:"tcpdomain"`
}
