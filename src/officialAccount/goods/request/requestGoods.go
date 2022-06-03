package request

type MainImage struct {
	URL string `json:"url"`
}

type ImageInfo struct {
	MainImageList []*MainImage `json:"main_image_list"`
}

type CategoryItem struct {
	CategoryName string `json:"category_name"`
}

type CategoryInfo struct {
	CategoryItem []*CategoryItem `json:"category_item"`
}

type OfficialCategoryInfo struct {
	CategoryItem []*CategoryItem `json:"category_item"`
}

type LinkInfo struct {
	URL      string `json:"url"`
	WxaAppID string `json:"wxa_appid"`
	LinkType string `json:"link_type"`
}

type ShopInfo struct {
	Source int `json:"source"`
}

type PriceInfo struct {
	MinPrice    int     `json:"min_price"`
	MaxPrice    float64 `json:"max_price"`
	MinOriPrice float64 `json:"min_ori_price"`
	MaxOriPrice float64 `json:"max_ori_price"`
}

type SaleInfo struct {
	SaleStatus string `json:"sale_status"`
	Stock      int    `json:"stock"`
}

type Custom struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type CustomInfo struct {
	CustomList []*Custom `json:"custom_list"`
}

type AttributeItem struct {
	AttributeKey   string `json:"attribute_key"`
	AttributeValue string `json:"attribute_value"`
}

type AttributeInfo struct {
	AttributeItem []*AttributeItem `json:"attribute_item"`
}

type SKUItem struct {
	SKUID         string         `json:"sku_id"`
	BarcodeType   string         `json:"barcode_type"`
	Barcode       int64          `json:"barcode"`
	ImageInfo     *ImageInfo     `json:"image_info"`
	LinkUrl       string         `json:"link_url"`
	PriceInfo     *PriceInfo     `json:"price_info"`
	SaleInfo      *SaleInfo      `json:"sale_info"`
	ShopInfo      *ShopInfo      `json:"shop_info"`
	AttributeInfo *AttributeInfo `json:"attribute_info,omitempty"`
}

type SKUInfo struct {
	SKUItem []*SKUItem `json:"sku_item"`
}

type TagItem struct {
	TagName string `json:"tag_name"`
}

type TagInfo struct {
	TagItem []*TagItem `json:"tag_item"`
}

type Product struct {
	PID                  string                `json:"pid"`
	ImageInfo            *ImageInfo            `json:"image_info"`
	CategoryInfo         *CategoryInfo         `json:"category_info"`
	OfficialCategoryInfo *OfficialCategoryInfo `json:"official_category_info"`
	LinkInfo             *LinkInfo             `json:"link_info"`
	Title                string                `json:"title"`
	SubTitle             string                `json:"sub_title"`
	Brand                string                `json:"brand"`
	ShopInfo             *ShopInfo             `json:"shop_info"`
	Desc                 string                `json:"desc"`
	PriceInfo            *PriceInfo            `json:"price_info"`
	SaleInfo             *SaleInfo             `json:"sale_info"`
	CustomInfo           *CustomInfo           `json:"custom_info,omitempty"`
	SKUInfo              *SKUInfo              `json:"sku_info"`
	PartialUpdate        int                   `json:"partial_update"`
	TagInfo              *TagInfo              `json:"tag_info,omitempty"`
}

type RequestProduct struct {
	Product []*Product `json:"product"`
}

// ----------------------------------------

type ProductID struct {
	PID string `json:"pid"`
}

type RequestProductGet struct {
	Product *ProductID `json:"product"`
}
