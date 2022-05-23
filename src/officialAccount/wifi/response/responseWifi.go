package response

import "github.com/ArtisanCloud/PowerWeChat/src/kernel/response"

type Wifi struct {
	ShopID         string `json:"shop_id"`
	StatisTime     int64  `json:"statis_time"`
	TotalUser      int    `json:"total_user"`
	HomepageUV     int    `json:"homepage_uv"`
	NewFans        int    `json:"new_fans"`
	TotalFans      int    `json:"total_fans"`
	WXConnectUser  int    `json:"wxconnect_user"`
	ConnectMsgUser int    `json:"connect_msg_user"`
}

type ResponseWifiSummary struct {
	*response.ResponseOfficialAccount

	Data []*Wifi `json:"data"`
}

// ------------------------------------------------------------

type QRCode struct {
	QrcodeURL string `json:"qrcode_url"`
}

type ResponseWifiQrCodeURL struct {
	*response.ResponseOfficialAccount

	Data *QRCode `json:"data"`
}

// ------------------------------------------------------------
