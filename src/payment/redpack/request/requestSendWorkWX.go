package request

type RequestSendWorkWX struct {

	NonceStr            string   `xml:"nonce_str"`
	Sign                string   `xml:"sign"`
	MchBillNO           string   `xml:"mch_billno"`
	MchID               string   `xml:"mch_id"`
	WXAppID             string   `xml:"wxappid"`
	SenderName          string   `xml:"sender_name"`
	SenderHeaderMediaID string   `xml:"sender_header_media_id"`
	ReOpenID            string   `xml:"re_openid"`
	TotalAmount         string   `xml:"total_amount"`
	Wishing             string   `xml:"wishing"`
	ActName             string   `xml:"act_name"`
	Remark              string   `xml:"remark"`
	WorkWXSign          string   `xml:"workwx_sign"`
}

