package request

import "encoding/xml"

type SendMiniProgramHB struct {
	XMLName     xml.Name `xml:"xml"`
	Text        string   `xml:",chardata"`
	ActName     string   `xml:"act_name"`
	MchBillno   string   `xml:"mch_billno"`
	MchID       string   `xml:"mch_id"`
	NonceStr    string   `xml:"nonce_str"`
	NotifyWay   string   `xml:"notify_way"`
	ReOpenid    string   `xml:"re_openid"`
	Remark      string   `xml:"remark"`
	SendName    string   `xml:"send_name"`
	TotalAmount int   `xml:"total_amount"`
	TotalNum    int   `xml:"total_num"`
	Wishing     string   `xml:"wishing"`
	Wxappid     string   `xml:"wxappid"`
	Sign        string   `xml:"sign"`
}

