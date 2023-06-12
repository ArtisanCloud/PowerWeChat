package request

type Media struct {
	Filename string `json:"filename"`
	Sha256   string `json:"sha256"`
}

type RequestMediaUpload struct {
	File  string `json:"file"`
	Media *Media `json:"media"`
}
