package request

type RequestSendMiniProgramNormal struct {
	MchBillNO   string `xml:"mch_billno" json:"mch_billno"`
	MchID       string `xml:"mch_id" json:"mch_id"`
	WXAppID     string `xml:"wxappid" json:"wxappid"`
	SendName    string `xml:"send_name" json:"send_name"`
	ReOpenID    string `xml:"re_openid" json:"re_openid"`
	TotalAmount int    `xml:"total_amount" json:"total_amount"`
	TotalNum    int    `xml:"total_num" json:"total_num"`
	Wishing     string `xml:"wishing" json:"wishing"`
	ActName     string `xml:"act_name" json:"act_name"`
	Remark      string `xml:"remark" json:"remark"`
	NotifyWay   string `xml:"notify_way" json:"notify_way"`
	SceneID     string `xml:"scene_id" json:"scene_id"`
}
