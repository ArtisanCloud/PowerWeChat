package request

type RequestSendRedPack struct {
	Sign        string   `xml:"sign" json:"sign"`
	MchBillno   string   `xml:"mch_billno" json:"mch_billno"`
	MchID       string   `xml:"mch_id" json:"mch_id"`
	WxappID     string   `xml:"wxappid" json:"wxappid"`
	SendName    string   `xml:"send_name" json:"send_name"`
	ReOpenID    string   `xml:"re_openid" json:"re_openid"`
	TotalAmount string   `xml:"total_amount" json:"total_amount"`
	TotalNum    string   `xml:"total_num" json:"total_num"`
	Wishing     string   `xml:"wishing" json:"wishing"`
	ClientIp    string   `xml:"client_ip" json:"client_ip"`
	ActName     string   `xml:"act_name" json:"act_name"`
	Remark      string   `xml:"remark" json:"remark"`
	SceneID     string   `xml:"scene_id" json:"scene_id"`
	NonceStr    string   `xml:"nonce_str" json:"nonce_str"`
	RiskInfo    string   `xml:"risk_info" json:"risk_info"`
}
