package request

type Struct struct {
	URL string `json:"url"`
}

type RequestWifiSetHomePage struct {
	ShopID     int     `json:"shop_id"`
	TemplateID int     `json:"template_id"`
	Struct     *Struct `json:"struct"`
}

// ------------------------------------------------------------

type RequestWifiCardSet struct {
	ShopID       int    `json:"shop_id"`
	CardID       string `json:"card_id"`
	CardDescribe string `json:"card_describe"`
	StartTime    int    `json:"start_time"`
	EndTime      int    `json:"end_time"`
}
