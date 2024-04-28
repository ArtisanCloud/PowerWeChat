package request

type RequestComplaints struct {
	Limit            int8   `json:"limit"`
	Offset           int8   `json:"offset"`
	BeginDate        string `json:"begin_date"`
	EndDate          string `json:"end_date"`
	ComplaintedMchId string `json:"complainted_mchid"`
}
