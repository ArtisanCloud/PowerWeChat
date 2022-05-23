package request

type Struct struct {
	URL string `json:"url"`
}

type RequestWifiSetHomePage struct {
	ShopID     int     `json:"shop_id"`
	TemplateID int     `json:"template_id"`
	Struct     *Struct `json:"struct"`
}
