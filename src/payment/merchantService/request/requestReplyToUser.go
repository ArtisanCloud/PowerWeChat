package request

type RequestReplyToUser struct {
	ComplaintedMchid string   `json:"complainted_mchid"`
	ResponseContent  string   `json:"response_content"`
	ResponseImages   []string `json:"response_images"`
	JumpUrl          string   `json:"jump_url"`
	JumpUrlText      string   `json:"jump_url_text"`
}
