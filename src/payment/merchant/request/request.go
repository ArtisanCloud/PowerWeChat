package request

type Meta struct {
	Filename string `json:"filename"`
	Sha256   string `json:"sha256"`
}

type RequestMediaUpload struct {
	File string `json:"file"`
	Meta *Meta  `json:"meta"`
}
