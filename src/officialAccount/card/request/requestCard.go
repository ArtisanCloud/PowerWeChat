package request

type DateInfo struct {
	Type           string `json:"type"`
	BeginTimestamp int    `json:"begin_timestamp"`
	EndTimestamp   int    `json:"end_timestamp"`
}

type SKU struct {
	Quantity int `json:"quantity"`
}

type BaseInfo struct {
	LogoUrl           string    `json:"logo_url"`
	BrandName         string    `json:"brand_name"`
	CodeType          string    `json:"code_type"`
	Title             string    `json:"title"`
	Color             string    `json:"color"`
	Notice            string    `json:"notice"`
	ServicePhone      string    `json:"service_phone"`
	Description       string    `json:"description"`
	DateInfo          *DateInfo `json:"date_info"`
	SKU               *SKU      `json:"sku"`
	UseLimit          int       `json:"use_limit"`
	GetLimit          int       `json:"get_limit"`
	UseCustomCode     bool      `json:"use_custom_code"`
	BindOpenid        bool      `json:"bind_openid"`
	CanShare          bool      `json:"can_share"`
	CanGiveFriend     bool      `json:"can_give_friend"`
	LocationIdList    []int     `json:"location_id_list"`
	CenterTitle       string    `json:"center_title"`
	CenterSubTitle    string    `json:"center_sub_title"`
	CenterUrl         string    `json:"center_url"`
	CustomUrlName     string    `json:"custom_url_name"`
	CustomUrl         string    `json:"custom_url"`
	CustomUrlSubTitle string    `json:"custom_url_sub_title"`
	PromotionUrlName  string    `json:"promotion_url_name"`
	PromotionUrl      string    `json:"promotion_url"`
	Source            string    `json:"source"`
}

type UseCondition struct {
	AcceptCategory          string `json:"accept_category"`
	RejectCategory          string `json:"reject_category"`
	CanUseWithOtherDiscount bool   `json:"can_use_with_other_discount"`
}

type Abstract struct {
	Abstract    string   `json:"abstract"`
	IconURLList []string `json:"icon_url_list"`
}

type TextImage struct {
	ImageURL string `json:"image_url"`
	Text     string `json:"text"`
}

type TimeLimit struct {
	Type        string `json:"type"`
	BeginHour   int    `json:"begin_hour,omitempty"`
	EndHour     int    `json:"end_hour,omitempty"`
	BeginMinute int    `json:"begin_minute,omitempty"`
	EndMinute   int    `json:"end_minute,omitempty"`
}

type AdvancedInfo struct {
	UseCondition    *UseCondition `json:"use_condition"`
	Abstract        *Abstract     `json:"abstract"`
	TextImageList   []*TextImage  `json:"text_image_list"`
	TimeLimit       []*TimeLimit  `json:"time_limit"`
	BusinessService []string      `json:"business_service"`
}

// ------------------------------------------------------------

type GroupOn struct {
	CardInterface
	BaseInfo     *BaseInfo     `json:"base_info"`
	AdvancedInfo *AdvancedInfo `json:"advanced_info"`
	DealDetail   string        `json:"deal_detail"`
}

type Cash struct {
	CardInterface
	BaseInfo     *BaseInfo     `json:"base_info"`
	AdvancedInfo *AdvancedInfo `json:"advanced_info"`
	LeastCost    int           `json:"least_cost"`
	ReduceCost   int           `json:"reduce_cost"`
}

type Discount struct {
	CardInterface
	BaseInfo     *BaseInfo     `json:"base_info"`
	AdvancedInfo *AdvancedInfo `json:"advanced_info"`
	Discount     int           `json:"discount"`
}

type Gift struct {
	CardInterface
	BaseInfo     *BaseInfo     `json:"base_info"`
	AdvancedInfo *AdvancedInfo `json:"advanced_info"`
	Gift         string        `json:"gift"`
}

type GeneralCoupon struct {
	CardInterface
	BaseInfo      *BaseInfo     `json:"base_info"`
	AdvancedInfo  *AdvancedInfo `json:"advanced_info"`
	DefaultDetail string        `json:"default_detail"`
}

type MemberCard struct {
	CardInterface
	BaseInfo     *BaseInfo `json:"base_info"`
	BonusCleared string    `json:"bonus_cleared"`
	BonusRules   string    `json:"bonus_rules"`
	Prerogative  string    `json:"prerogative"`
}

type ScenicTicket struct {
	CardInterface
	BaseInfo *BaseInfo `json:"base_info"`
}
type MovieTicket struct {
	CardInterface
	BaseInfo *BaseInfo `json:"base_info"`
}

type BoardingPass struct {
	CardInterface
	BaseInfo *BaseInfo `json:"base_info"`
}

type MeetingTicket struct {
	CardInterface
	BaseInfo *BaseInfo `json:"base_info"`
}

type BusTicket struct {
	CardInterface
	BaseInfo *BaseInfo `json:"base_info"`
}

// -----------------------------------------------------------------

type CardInterface interface {
	GetCardType() string
}

type Card struct {
	CardType      string         `json:"card_type"`
	Groupon       *GroupOn       `json:"groupon"`
	Cash          *Cash          `json:"cash"`
	Discount      *Discount      `json:"discount"`
	Gift          *Gift          `json:"gift"`
	GeneralCoupon *GeneralCoupon `json:"general_coupon"`
	MemberCard    *MemberCard    `json:"member_card"`
	ScenicTicket  *ScenicTicket  `json:"scenic_ticket"`
	MovieTicket   *MovieTicket   `json:"movie_ticket"`
	BoardingPass  *BoardingPass  `json:"boarding_pass"`
	MeetingTicket *MeetingTicket `json:"meeting_ticket"`
	BusTicket     *BusTicket     `json:"bus_ticket"`
}

type RequestCardCreate struct {
	Card *Card `json:"card"`
}
