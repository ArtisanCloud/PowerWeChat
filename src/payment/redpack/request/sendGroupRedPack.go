package request

import "encoding/xml"

type SendGroupRedPack struct {
	XMLName     xml.Name `xml:"xml"`
	Text        string   `xml:",chardata"`
	Sign        string   `xml:"sign"`
	MchBillno   string   `xml:"mch_billno"`
	MchID       string   `xml:"mch_id"`
	Wxappid     string   `xml:"wxappid"`
	SendName    string   `xml:"send_name"`
	ReOpenid    string   `xml:"re_openid"`
	TotalAmount string   `xml:"total_amount"`
	AmtType     string   `xml:"amt_type"`
	TotalNum    string   `xml:"total_num"`
	Wishing     string   `xml:"wishing"`
	ActName     string   `xml:"act_name"`
	Remark      string   `xml:"remark"`
	SceneID     string   `xml:"scene_id"`
	NonceStr    string   `xml:"nonce_str"`
	RiskInfo    string   `xml:"risk_info"`
}

