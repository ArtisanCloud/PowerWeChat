package request

type RequestPayTransferToPocket struct {

	AppID          string   `xml:",appid"`
	MchID          string   `xml:"mch_id"`
	DeviceInfo     string   `xml:"device_info"`
	NonceStr       string   `xml:"nonce_str"`
	Sign           string   `xml:"sign"`
	PartnerTradeNO string   `xml:"partner_trade_no"`
	OpenID         string   `xml:"openid"`
	CheckName      string   `xml:"check_name"`
	ReUserName     string   `xml:"re_user_name"`
	Amount         string   `xml:"amount"`
	Desc           string   `xml:"desc"`
	SPBillCreateIP string   `xml:"spbill_create_ip"`
	WorkWXSign     string   `xml:"workwx_sign"`
	WwMsgType      string   `xml:"ww_msg_type"`
	ActName        string   `xml:"act_name"`
}
