package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

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
	response.ResponseOfficialAccount

	Data []*Wifi `json:"data"`
}

// ------------------------------------------------------------

type QRCode struct {
	QrcodeURL string `json:"qrcode_url"`
}

type ResponseWifiQrCodeURL struct {
	response.ResponseOfficialAccount

	Data *QRCode `json:"data"`
}

// ------------------------------------------------------------

type Device struct {
	ShopID       int    `json:"shop_id"`
	CardStatus   int    `json:"card_status"`
	CardID       string `json:"card_id"`
	CardDescribe string `json:"card_describe"`
	StartDate    int    `json:"start_date"`
	EndDate      int    `json:"end_date"`
}

type ResponseWifiCardGet struct {
	response.ResponseOfficialAccount

	Data *Device `json:"data"`
}

// ------------------------------------------------------------

type Record struct {
	ShopID       int      `json:"shop_id"`
	SSID         string   `json:"ssid"`
	BssID        string   `json:"bssid"`
	ProtocolType int      `json:"protocol_type"`
	ShopName     string   `json:"shop_name"`
	SSIDList     []string `json:"ssid_list"`
	SID          string   `json:"sid"`
	PoiID        string   `json:"poi_id"`
}

type DeviceList struct {
	TotalCount int       `json:"totalcount"`
	PageIndex  int       `json:"pageindex"`
	PageCount  int       `json:"pagecount"`
	Records    []*Record `json:"records"`
}

type ResponseWifiDeviceList struct {
	response.ResponseOfficialAccount

	Data *DeviceList `json:"data"`
}

// ------------------------------------------------------------

type SSIDPassword struct {
	SSID     string `json:"ssid"`
	Password string `json:"password"`
}

type Shop struct {
	ShopName              string          `json:"shop_name"`
	Ssid                  string          `json:"ssid"`
	SsidList              []string        `json:"ssid_list"`
	SsidPasswordList      []*SSIDPassword `json:"ssid_password_list"`
	Password              string          `json:"password"`
	ProtocolType          int             `json:"protocol_type"`
	ApCount               int             `json:"ap_count"`
	TemplateId            int             `json:"template_id"`
	HomepageUrl           string          `json:"homepage_url"`
	BarType               int             `json:"bar_type"`
	Sid                   string          `json:"sid"`
	PoiId                 string          `json:"poi_id"`
	HomepageWxaUserName   string          `json:"homepage_wxa_user_name"`
	HomepageWxaPath       string          `json:"homepage_wxa_path"`
	FinishPageURL         string          `json:"finishpage_url"`
	FinishPageWXAUserName string          `json:"finishpage_wxa_user_name"`
	FinishPageWXAPath     string          `json:"finishpage_wxa_path"`
	FinishPageType        int             `json:"finishpage_type"`
}

type ResponseWifiShopGet struct {
	response.ResponseOfficialAccount

	Data *Shop `json:"data"`
}

// ------------------------------------------------------------

type ShopList struct {
	TotalCount int       `json:"totalcount"`
	PageIndex  int       `json:"pageindex"`
	PageCount  int       `json:"pagecount"`
	Records    []*Record `json:"records"`
}

type ResponseWifiShopList struct {
	response.ResponseOfficialAccount

	Data *ShopList `json:"data"`
}
