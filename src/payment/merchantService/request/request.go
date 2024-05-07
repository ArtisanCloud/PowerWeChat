package request

type Meta struct {
	Filename string `json:"filename"`
	Sha256   string `json:"sha256"`
}

type RequestMediaUpload struct {
	File string `json:"file"`
	Meta *Meta  `json:"meta"`
}

type RequestComplaints struct {
	Limit            int8   `json:"limit"`
	Offset           int8   `json:"offset"`
	BeginDate        string `json:"begin_date"`
	EndDate          string `json:"end_date"`
	ComplaintedMchId string `json:"complainted_mchid"`
}
