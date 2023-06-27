package request

type DateInfo struct {
	Type           string `json:"type"`
	BeginTimestamp int    `json:"begin_timestamp"`
	EndTimestamp   int    `json:"end_timestamp"`
}

type SKU struct {
	Quantity int `json:"quantity"`
}
type PayInfo struct {
	SwipeCard *SwipeCard `json:"swipe_card,omitempty"`
}
type SwipeCard struct {
	IsSwapCard bool `json:"is_swipe_card"`
}
type BaseInfo struct {
	SubMerchantInfo   *SubMerchantID `json:"sub_merchant_info,omitempty"`
	LogoUrl           string         `json:"logo_url,omitempty"`
	BrandName         string         `json:"brand_name,omitempty"`
	CodeType          string         `json:"code_type,omitempty"`
	Title             string         `json:"title,omitempty"`
	Color             string         `json:"color,omitempty"`
	Notice            string         `json:"notice,omitempty"`
	ServicePhone      string         `json:"service_phone,omitempty"`
	Description       string         `json:"description,omitempty"`
	DateInfo          *DateInfo      `json:"date_info,omitempty"`
	SKU               *SKU           `json:"sku,omitempty"`
	UseLimit          int            `json:"use_limit,omitempty"`
	GetLimit          int            `json:"get_limit,omitempty"`
	UseCustomCode     bool           `json:"use_custom_code,omitempty"`
	BindOpenid        bool           `json:"bind_openid,omitempty"`
	CanShare          bool           `json:"can_share,omitempty"`
	CanGiveFriend     bool           `json:"can_give_friend,omitempty"`
	LocationIdList    []int          `json:"location_id_list,omitempty"`
	CenterTitle       string         `json:"center_title,omitempty"`
	CenterSubTitle    string         `json:"center_sub_title,omitempty"`
	CenterUrl         string         `json:"center_url,omitempty"`
	CustomUrlName     string         `json:"custom_url_name,omitempty"`
	CustomUrl         string         `json:"custom_url,omitempty"`
	CustomUrlSubTitle string         `json:"custom_url_sub_title,omitempty"`
	PromotionUrlName  string         `json:"promotion_url_name,omitempty"`
	PromotionUrl      string         `json:"promotion_url,omitempty"`
	Source            string         `json:"source,omitempty"`
	PayInfo           *PayInfo       `json:"pay_info,omitempty"`
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
	// CardInterface
	BaseInfo     *BaseInfo     `json:"base_info"`
	AdvancedInfo *AdvancedInfo `json:"advanced_info"`
	DealDetail   string        `json:"deal_detail"`
}

type Cash struct {
	// CardInterface
	BaseInfo     *BaseInfo     `json:"base_info"`
	AdvancedInfo *AdvancedInfo `json:"advanced_info"`
	LeastCost    int           `json:"least_cost"`
	ReduceCost   int           `json:"reduce_cost"`
}

type Discount struct {
	// CardInterface
	BaseInfo     *BaseInfo     `json:"base_info"`
	AdvancedInfo *AdvancedInfo `json:"advanced_info"`
	Discount     int           `json:"discount"`
}

type Gift struct {
	// CardInterface
	BaseInfo     *BaseInfo     `json:"base_info"`
	AdvancedInfo *AdvancedInfo `json:"advanced_info"`
	Gift         string        `json:"gift"`
}

type GeneralCoupon struct {
	// CardInterface
	BaseInfo      *BaseInfo     `json:"base_info"`
	AdvancedInfo  *AdvancedInfo `json:"advanced_info"`
	DefaultDetail string        `json:"default_detail"`
}

type MemberCard struct {
	// CardInterface
	BackgroundPicURL string     `json:"background_pic_url,omitempty"`
	BaseInfo         *BaseInfo  `json:"base_info"`
	Prerogative      string     `json:"prerogative,omitempty"`
	BonusCleared     string     `json:"bonus_cleared,omitempty"`
	SupplyBonus      bool       `json:"supply_bonus,omitempty"`
	BonusURL         string     `json:"bonus_url,omitempty"`
	SupplyBalance    bool       `json:"supply_balance,omitempty"`
	BalanceRules     string     `json:"balance_rules,omitempty"`
	BalanceURL       string     `json:"balance_url,omitempty"`
	BonusRules       string     `json:"bonus_rules,omitempty"`
	BonusRule        *BonusRule `json:"bonus_rule,omitempty"`
	Discount         int        `json:"discount,omitempty"`
	ActivateURL      string     `json:"activate_url,omitempty"`
	WXActivate       bool       `json:"wx_activate,omitempty"`
	AutoActivate     bool       `json:"auto_activate,omitempty"`
}

func (m *MemberCard) GetCardType() string {
	return "member_card"
}

type BonusRule struct {
	CostMoneyUnit        int `json:"cost_money_unit,omitempty"`
	IncreaseBonus        int `json:"increase_bonus,omitempty"`
	MaxIncreaseBonus     int `json:"max_increase_bonus,omitempty"`
	InitIncreaseBonus    int `json:"init_increase_bonus,omitempty"`
	CostBonusUnit        int `json:"cost_bonus_unit,omitempty"`
	ReduceMoney          int `json:"reduce_money,omitempty"`
	LeastMoneyToUseBonus int `json:"least_money_to_use_bonus,omitempty"`
	MaxReduceBonus       int `json:"max_reduce_bonus,omitempty"`
}
type ScenicTicket struct {
	// CardInterface
	BaseInfo *BaseInfo `json:"base_info"`
}
type MovieTicket struct {
	// CardInterface
	BaseInfo *BaseInfo `json:"base_info"`
}

type BoardingPass struct {
	// CardInterface
	BaseInfo *BaseInfo `json:"base_info"`
}

type MeetingTicket struct {
	// CardInterface
	BaseInfo *BaseInfo `json:"base_info"`
}

type BusTicket struct {
	// CardInterface
	BaseInfo *BaseInfo `json:"base_info"`
}

// -----------------------------------------------------------------

type CardInterface interface {
	GetCardType() string
}

type Card struct {
	CardType      string         `json:"card_type"`
	Groupon       *GroupOn       `json:"groupon,omitempty"`
	Cash          *Cash          `json:"cash,omitempty"`
	Discount      *Discount      `json:"discount,omitempty"`
	Gift          *Gift          `json:"gift,omitempty"`
	GeneralCoupon *GeneralCoupon `json:"general_coupon,omitempty"`
	MemberCard    *MemberCard    `json:"member_card,omitempty"`
	ScenicTicket  *ScenicTicket  `json:"scenic_ticket,omitempty"`
	MovieTicket   *MovieTicket   `json:"movie_ticket,omitempty"`
	BoardingPass  *BoardingPass  `json:"boarding_pass,omitempty"`
	MeetingTicket *MeetingTicket `json:"meeting_ticket,omitempty"`
	BusTicket     *BusTicket     `json:"bus_ticket,omitempty"`
}

func (c Card) GetCardType() string {
	return c.CardType
}

type RequestCardCreate struct {
	Card *Card `json:"card"`
}
