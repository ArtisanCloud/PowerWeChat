package request

type RequestSendGroupRedPack struct {

	Sign        string   `xml:"sign"`
	MchBillNO   string   `xml:"mch_billno"`
	MchID       string   `xml:"mch_id"`
	WXAppID     string   `xml:"wxappid"`
	SendName    string   `xml:"send_name"`
	ReOpenID    string   `xml:"re_openid"`
	TotalAmount int64    `xml:"total_amount"`
	TotalNum    int32    `xml:"total_num"`
	AmtType     string   `xml:"amt_type"`
	Wishing     string   `xml:"wishing"`
	ActName     string   `xml:"act_name"`
	Remark      string   `xml:"remark"`
	SceneID     string   `xml:"scene_id"`
	NonceStr    string   `xml:"nonce_str"`
	RiskInfo    string   `xml:"risk_info"`
}
